package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/earn-alliance/wallet-commander-cli/pkg/abi"
	"github.com/earn-alliance/wallet-commander-cli/pkg/constants"
	"github.com/earn-alliance/wallet-commander-cli/pkg/store"
	pkgTypes "github.com/earn-alliance/wallet-commander-cli/pkg/types"
	"github.com/earn-alliance/wallet-commander-cli/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/patrickmn/go-cache"
)

const (
	AXIE_INFINITY_GRAPHQL_GATEWAY = "https://graphql-gateway.axieinfinity.com/graphql"
	HOUR_PER_DAY                  = 24
	DAYS_BETWEEN_CLAIM            = 14
	DURATION_BETWEEN_CLAIMS       = time.Hour * HOUR_PER_DAY * DAYS_BETWEEN_CLAIM
)

var (
	// TODO: Dynamically determine this value based on free transfers
	defaultGasPriceWei = big.NewInt(1 * params.GWei)
)

type Client interface {
	TransferSlp(ctx context.Context, privateKey, to string, amount int) (string, error)
	TransferAxie(ctx context.Context, privateKey, to string, axieId int) (string, error)
	GetTransactionReceipt(ctx context.Context, txHash string) (*types.Receipt, error)
	GetRoninWalletBalance(ctx context.Context, tokenTypeAddress, address string) (float64, error)
	GetClaimableAmount(ctx context.Context, address string) (*ClaimableResponse, error)
	ClaimSlp(ctx context.Context, address, privateKeyStr string) (string, error)
	ClaimOriginSlp(ctx context.Context, address, privateKeyStr string) (string, error)
	GetWalletTransactionHistory(address string) (*pkgTypes.WalletTransactionHistory, error)
	GetClaimPayload(ctx context.Context, address, privateKeyStr string) (*ClaimableResponse, error)
	GetClaimOriginPayload(ctx context.Context, address, privateKeyStr string) (*OriginClaimableResponse, error)
	CreateAccessToken(address, privateKeyStr string) (string, error)
}

type AxieClient struct {
	// Used for submitting transactions on ronin blockchain (writes)
	freeEthClient *ethclient.Client
	// Used for querying ronin blockchain (reads)
	ethClient *ethclient.Client

	slpClient         *abi.Slp
	portalClient      *abi.Portal
	axieClient        *abi.Axie
	marketplaceClient *abi.Marketplace

	jwtStore   store.JwtStore
	nonceCache *cache.Cache
	mu         sync.RWMutex
}

func (c *AxieClient) GetWalletTransactionHistory(address string) (*pkgTypes.WalletTransactionHistory, error) {
	const transactionQueryUrl = "https://explorer-api.roninchain.com/tokentxs?addr=%s&from=%d&size=%d&token=ERC20"
	const transactionsPerPage = 100
	var walletTransactionHistory = new(pkgTypes.WalletTransactionHistory)
	startFrom := 0

	for {
		url := fmt.Sprintf(
			transactionQueryUrl,
			utils.RoninAddrToEthAddr(address),
			startFrom,
			transactionsPerPage,
		)

		respBytes, err := utils.CallGetHttpApi(url, utils.DefaultAxieSiteRequestHeaders)

		if err != nil {
			return nil, err
		}

		var resp pkgTypes.WalletTransactionHistory

		if err := json.Unmarshal(respBytes, &resp); err != nil {
			return nil, err
		}

		for i, result := range resp.Results {
			result.TimestampStr = time.Unix(result.Timestamp, 0).Format("2006-01-02 15:04:05")
			resp.Results[i] = result
		}

		if walletTransactionHistory == nil {
			walletTransactionHistory = &resp
		} else {
			walletTransactionHistory.Results = append(walletTransactionHistory.Results, resp.Results...)
			walletTransactionHistory.Total += resp.Total
		}

		if resp.Total == transactionsPerPage {
			startFrom += transactionsPerPage
		} else {
			break
		}
	}

	return walletTransactionHistory, nil
}

func createRpcClient(url string) (*rpc.Client, error) {
	rpcClient, err := rpc.DialHTTP(url)

	if err != nil {
		return nil, err
	}

	rpcClient.SetHeader("context-type", "application/json")
	rpcClient.SetHeader("user-agent", "wallet-commander")

	return rpcClient, nil
}

func New() (Client, error) {
	freeRpcClient, err := createRpcClient(constants.RONIN_PROVIDER_FREE_URI)

	if err != nil {
		return nil, err
	}

	rpcClient, err := createRpcClient(constants.RONIN_PROVIDER_RPC_URI)

	if err != nil {
		return nil, err
	}

	freeEthClient := ethclient.NewClient(freeRpcClient)
	ethClient := ethclient.NewClient(rpcClient)

	slpClient, err := abi.NewSlp(common.HexToAddress(constants.SLP_CONTRACT), freeEthClient)

	if err != nil {
		return nil, err
	}

	axieClient, err := abi.NewAxie(common.HexToAddress(constants.AXIE_CONTRACT), freeEthClient)

	if err != nil {
		return nil, err
	}

	marketplaceClient, err := abi.NewMarketplace(common.HexToAddress(constants.MARKETPLACE_CONTRACT), freeEthClient)

	if err != nil {
		return nil, err
	}

	portalClient, err := abi.NewPortal(common.HexToAddress(constants.PORTAL_CONTRACT), freeEthClient)

	if err != nil {
		return nil, err
	}

	nonceCache := cache.New(15*time.Minute, 20*time.Minute)

	return &AxieClient{
		freeEthClient:     freeEthClient,
		ethClient:         ethClient,
		slpClient:         slpClient,
		axieClient:        axieClient,
		marketplaceClient: marketplaceClient,
		portalClient:      portalClient,
		nonceCache:        nonceCache,
	}, nil
}

func (c *AxieClient) createTransactOps(ctx context.Context, privateKeyStr string) (*bind.TransactOpts, error) {
	return c.createTransactOpsWithGasLimitRange(ctx, privateKeyStr, 1000000)
}

func (c *AxieClient) createTransactOpsWithGasLimitRange(ctx context.Context, privateKeyStr string, gasLimit uint64) (*bind.TransactOpts, error) {
	privateKey, err := crypto.ToECDSA(common.FromHex(privateKeyStr))

	if err != nil {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	c.mu.Lock()
	defer c.mu.Unlock()

	nonce, err := c.nonceAt(ctx, fromAddress)

	if err != nil {
		return nil, err
	}

	_ = c.incrNonce(ctx, fromAddress)

	ops, err := bind.NewKeyedTransactorWithChainID(privateKey, constants.RONIN_CHAIN_ID)

	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	// gasLimit := uint64(rand.Int63n(upperGasLimit-lowerGasLimit+1) + lowerGasLimit)
	ops.Nonce = big.NewInt(int64(nonce))
	ops.GasPrice = defaultGasPriceWei
	ops.GasLimit = gasLimit
	ops.Context = ctx

	ops.Signer = func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(transaction, types.NewEIP155Signer(constants.RONIN_CHAIN_ID), privateKey)
	}

	return ops, nil
}

func (c *AxieClient) nonceAt(ctx context.Context, address common.Address) (uint64, error) {
	nonce, err := c.freeEthClient.NonceAt(ctx, address, nil)
	if err != nil {
		return 0, err
	}

	val, found := c.nonceCache.Get(address.String())
	if !found {
		return nonce, nil
	}

	cache := val.(uint64)

	if nonce > cache {
		return nonce, nil
	} else {
		return cache, nil
	}
}

func (c *AxieClient) incrNonce(ctx context.Context, address common.Address) error {
	nonce, err := c.nonceAt(ctx, address)
	if err != nil {
		return err
	}

	c.nonceCache.Set(address.String(), nonce+1, cache.DefaultExpiration)
	return nil
}

func (c *AxieClient) TransferSlp(ctx context.Context, privateKey, to string, amount int) (string, error) {
	ops, err := c.createTransactOps(ctx, privateKey)

	if err != nil {
		return "", err
	}

	tx, err := c.slpClient.Transfer(ops, common.HexToAddress(to), big.NewInt(int64(amount)))

	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

func (c *AxieClient) TransferAxie(ctx context.Context, privateKey, to string, axieId int) (string, error) {
	ops, err := c.createTransactOps(ctx, privateKey)

	if err != nil {
		return "", err
	}

	tx, err := c.axieClient.SafeTransferFrom(ops, ops.From, common.HexToAddress(utils.RoninAddrToEthAddr(to)), big.NewInt(int64(axieId)))

	if err != nil {
		return "", err
	}

	return tx.Hash().String(), err
}

// TODO: Test
func (c *AxieClient) BreedAxie(ctx context.Context, privateKey string, dadAxieId, momAxieId int) (string, error) {
	ops, err := c.createTransactOps(ctx, privateKey)

	if err != nil {
		return "", err
	}

	tx, err := c.axieClient.BreedAxies(ops, big.NewInt(int64(dadAxieId)), big.NewInt(int64(momAxieId)))

	if err != nil {
		return "", err
	}

	return tx.Hash().String(), err
}

func (c *AxieClient) GetTransactionReceipt(ctx context.Context, txHash string) (*types.Receipt, error) {
	// All writes go to freeEthClient, so we only need to check transaction receipt here
	return c.freeEthClient.TransactionReceipt(ctx, common.HexToHash(txHash))
}

func (c *AxieClient) GetRoninWalletBalance(ctx context.Context, tokenTypeAddress, address string) (float64, error) {
	// TODO: Cache clients
	client, err := abi.NewRoninBalance(common.HexToAddress(tokenTypeAddress), c.ethClient)

	if err != nil {
		return 0, err
	}

	balance, err := client.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, common.HexToAddress(address))

	if err != nil {
		return 0, err
	}

	value := float64(balance.Int64())

	if tokenTypeAddress == constants.WETH_CONTRACT {
		value /= 1000000000000000000 // Convert to decimal from wei
	}

	return value, nil
}

type T struct {
}
type ClaimableResponse struct {
	ClientID          string `json:"clientID"`
	ItemID            int    `json:"itemID"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	ImageUrl          string `json:"imageUrl"`
	Total             int    `json:"total"`
	BlockchainRelated struct {
		Signature struct {
			Signature string `json:"signature"`
			Amount    int    `json:"amount"`
			Timestamp int    `json:"timestamp"`
		} `json:"signature"`
		Balance     int `json:"balance"`
		Checkpoint  int `json:"checkpoint"`
		BlockNumber int `json:"blockNumber"`
	} `json:"blockchainRelated"`
	ClaimableTotal    int `json:"claimableTotal"`
	LastClaimedItemAt int `json:"lastClaimedItemAt"`
	RawTotal          int `json:"rawTotal"`
	RawClaimableTotal int `json:"rawClaimableTotal"`
}

// The following functions are so funny, it makes me think im not sure whats going on OR the devs suck
// amount that can be claimed as of today if CanClaim is true
func (c ClaimableResponse) GetClaimableAmount() int {
	if c.BlockchainRelated.Signature.Amount > c.BlockchainRelated.Checkpoint {
		return c.BlockchainRelated.Signature.Amount - c.BlockchainRelated.Checkpoint
	} else {
		return c.RawTotal - c.RawClaimableTotal
	}
}

func (c ClaimableResponse) CanClaim() bool {
	// Cannot claim until 14 days later
	return c.BlockchainRelated.Signature.Amount > c.BlockchainRelated.Checkpoint ||
		(c.LastClaimedItemAt > 0 && time.Unix(int64(c.LastClaimedItemAt), 0).Before(time.Now().Add(-1*DURATION_BETWEEN_CLAIMS)))
}

func (c ClaimableResponse) HoursToNextClaim() float64 {
	return time.Unix(int64(c.LastClaimedItemAt), 0).Add(DURATION_BETWEEN_CLAIMS).Sub(time.Now()).Hours()
}

func (c *AxieClient) GetClaimableAmount(ctx context.Context, address string) (*ClaimableResponse, error) {
	var resp ClaimableResponse
	respBytes, err := utils.CallGetHttpApi(fmt.Sprintf("https://game-api-pre.skymavis.com/v1/players/%s/items/1",
		utils.RoninAddrToEthAddr(address),
	), nil)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}

	//log.Printf("claim resp %v %s", resp, utils.RoninAddrToEthAddr(address))

	return &resp, nil
}

type GraphqlRequest struct {
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
	Query         string                 `json:"query"`
}

func (c *AxieClient) GetClaimPayload(ctx context.Context, address, privateKeyStr string) (*ClaimableResponse, error) {
	randomMsg, err := createRandomMessage()

	if err != nil {
		return nil, err
	}

	token, err := c.getJwtAccessToken(randomMsg, address, privateKeyStr)

	if err != nil {
		return nil, err
	}

	var headers = utils.DefaultAxieSiteRequestHeaders
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	var resp ClaimableResponse
	respBytes, err := utils.CallPostHttpApi(
		"https://game-api-pre.skymavis.com/v1/players/me/items/1/claim",
		nil,
		headers,
	)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}

	//log.Printf("claim resp %v", resp)

	return &resp, nil
}

type OriginClaimableResponse struct {
	ClientID     string `json:"clientID"`
	ItemID       string `json:"itemID"`
	Withdrawable int    `json:"withdrawable"`
	Quantity     int    `json:"quantity"`
	Item         struct {
		Category      string `json:"category"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		Rarity        string `json:"rarity"`
		TokenStandard string `json:"tokenStandard"`
		TokenId       string `json:"tokenId"`
	} `json:"item"`
}

type OriginWithdrawItemsResponse struct {
	UserAddress string `json:"userAddress"`
	Nonce       int64  `json:"nonce"`
	ExpiredAt   int64  `json:"expiredAt"`
	Signature   string `json:"signature"`
	Items       []struct {
		ItemId        string `json:"itemId"`
		TokenStandard string `json:"tokenStandard"`
		TokenAddress  string `json:"tokenAddress"`
		TokenId       string `json:"tokenId"`
		SignedAmount  string `json:"signedAmount"`
	} `json:"items"`
	ExtraData string `json:"extraData"`
}

func (c OriginClaimableResponse) CanClaim() bool {
	return c.Withdrawable > 0
}

func (c OriginClaimableResponse) GetClaimableAmount() int {
	return c.Withdrawable
}

func (c *AxieClient) GetClaimOriginPayload(ctx context.Context, address, privateKeyStr string) (*OriginClaimableResponse, error) {
	token, err := c.CreateAccessToken(address, privateKeyStr)
	if err != nil {
		return nil, err
	}
	return c.getClaimableSlpPayload(token)
}

func (c *AxieClient) getClaimableSlpPayload(token string) (*OriginClaimableResponse, error) {
	var headers = utils.DefaultAxieSiteRequestHeaders
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	var resp OriginClaimableResponse
	respBytes, err := utils.CallGetHttpApi(
		"https://game-api-origin.skymavis.com/v2/users/me/items/marketplace/slp",
		headers,
	)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *AxieClient) getWithdrawSlpPayload(token string, amount int) (*OriginWithdrawItemsResponse, error) {
	var headers = utils.DefaultAxieSiteRequestHeaders
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	var resp OriginWithdrawItemsResponse

	respBytes, err := utils.CallPostHttpApi(
		"https://game-api-origin.skymavis.com/v2/rpc/items/withdraw", map[string][]map[string]interface{}{
			"items": {
				{
					"amount": amount,
					"itemId": "slp",
				},
			},
		},
		headers,
	)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *AxieClient) ClaimOriginSlp(
	ctx context.Context, address, privateKeyStr string) (string, error) {

	token, err := c.CreateAccessToken(address, privateKeyStr)
	if err != nil {
		return "", err
	}

	claimableResponse, err := c.getClaimableSlpPayload(token)
	log.Printf("claimable resp %v", claimableResponse)
	if err != nil {
		return "", err
	}

	if !claimableResponse.CanClaim() {
		return "", errors.New(fmt.Sprintf(
			"No claimable slp",
		))
	}

	withdrawSlpResponse, err := c.getWithdrawSlpPayload(token, claimableResponse.GetClaimableAmount())

	//log.Printf("withdraw slp resp %v", withdrawSlpResponse)

	if err != nil {
		return "", errors.New(fmt.Sprintf(
			"Another transaction in progress. Please try again later",
		))
	}

	ops, err := c.createTransactOps(ctx, privateKeyStr)

	if err != nil {
		log.Printf("gas limit error %v", err)
		return "", err
	}

	tokenId, err := strconv.ParseInt(withdrawSlpResponse.Items[0].TokenId, 10, 32)
	if err != nil {
		return "", err
	}

	quantity, err := strconv.ParseInt(withdrawSlpResponse.Items[0].SignedAmount, 10, 32)
	if err != nil {
		return "", err
	}

	tx, err := c.portalClient.Withdraw(
		ops,
		abi.Struct2{
			Owner:     common.HexToAddress(withdrawSlpResponse.UserAddress),
			Nonce:     big.NewInt(withdrawSlpResponse.Nonce),
			ExpiredAt: big.NewInt(withdrawSlpResponse.ExpiredAt),
			Assets: []abi.Struct0{
				{
					Erc:      0,
					Addr:     common.HexToAddress(constants.SLP_CONTRACT),
					Id:       big.NewInt(tokenId),
					Quantity: big.NewInt(quantity),
				},
			},
			ExtraData: common.FromHex(withdrawSlpResponse.ExtraData),
		},
		common.FromHex(withdrawSlpResponse.Signature),
		[]common.Address{
			common.HexToAddress(constants.AXS_CONTRACT),
			common.HexToAddress(constants.WETH_CONTRACT),
			common.HexToAddress(constants.USDC_CONTRACT),
		},
	)

	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

func (c *AxieClient) ClaimSlp(
	ctx context.Context, address, privateKeyStr string) (string, error) {

	claimResponse, err := c.GetClaimPayload(ctx, address, privateKeyStr)

	if err != nil {
		return "", err
	}

	if !claimResponse.CanClaim() {
		return "", errors.New(fmt.Sprintf(
			"cannot claim. hours to next claim %f", claimResponse.HoursToNextClaim(),
		))
	}

	ops, err := c.createTransactOps(ctx, privateKeyStr)

	if err != nil {
		return "", err
	}

	tx, err := c.slpClient.Checkpoint(
		ops,
		common.HexToAddress(address),
		big.NewInt(int64(claimResponse.BlockchainRelated.Signature.Amount)),
		big.NewInt(int64(claimResponse.BlockchainRelated.Signature.Timestamp)),
		common.FromHex(claimResponse.BlockchainRelated.Signature.Signature),
	)

	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

func (c *AxieClient) CreateRandomMessage() (string, error) {
	respBytes, err := utils.CallPostHttpApi(AXIE_INFINITY_GRAPHQL_GATEWAY, GraphqlRequest{
		OperationName: "CreateRandomMessage",
		Variables:     map[string]interface{}{},
		Query: `mutation CreateRandomMessage {
			createRandomMessage
		}`,
	}, utils.DefaultAxieSiteRequestHeaders)

	if err != nil {
		return "", err
	}
	type CreateRandomMessageResponse struct {
		Data struct {
			RandomMessage string `json:"createRandomMessage"`
		} `json:"data"`
	}
	var response CreateRandomMessageResponse
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return "", err
	}

	if response.Data.RandomMessage == "" {
		return "", errors.New("random message empty")
	}
	return response.Data.RandomMessage, nil

}

func (c *AxieClient) CreateAccessToken(address, privateKeyStr string) (string, error) {
	randomMsg, err := c.CreateRandomMessage()

	if err != nil {
		return "", err
	}

	return c.getJwtAccessToken(randomMsg, address, privateKeyStr)
}

func (c *AxieClient) getJwtAccessToken(randomMsg, address, privateKeyStr string) (string, error) {
	if cachedToken := c.jwtStore.GetValidJwt(address); cachedToken != "" {
		log.Printf("Returning cached JWT token for usage")
		return cachedToken, nil
	}

	type CreateAccessTokenWithSignatureResponse struct {
		Data struct {
			CreateAccessTokenWithSignature struct {
				AccessToken string `json:"accessToken"`
			} `json:"createAccessTokenWithSignature"`
		} `json:"data"`
	}

	privateKey := crypto.ToECDSAUnsafe(common.FromHex(privateKeyStr))
	hash := utils.NodejsHashData([]byte(randomMsg))
	signedMsgBytes, err := crypto.Sign(hash, privateKey)

	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not sign message with err %v", err))
	}

	respBytes, err := utils.CallPostHttpApi(AXIE_INFINITY_GRAPHQL_GATEWAY, GraphqlRequest{
		OperationName: "CreateAccessTokenWithSignature",
		Variables: map[string]interface{}{
			"input": map[string]string{
				"mainnet":   "ronin",
				"owner":     utils.RoninAddrToEthAddr(address),
				"message":   randomMsg,
				"signature": hexutil.Encode(signedMsgBytes),
			},
		},
		Query: `mutation CreateAccessTokenWithSignature($input: SignatureInput!) {
			createAccessTokenWithSignature(input: $input) {
				newAccount
				result
				accessToken
				__typename
			}
		}`,
	}, utils.DefaultAxieSiteRequestHeaders)

	if err != nil {
		return "", err
	}

	var response CreateAccessTokenWithSignatureResponse
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return "", err
	}

	if response.Data.CreateAccessTokenWithSignature.AccessToken == "" {
		return "", errors.New("access token in response was empty")
	}

	token := response.Data.CreateAccessTokenWithSignature.AccessToken
	//log.Printf("Created got access token %s", token)

	// Cache token for future usage
	c.jwtStore.StoreJwtToken(address, token)

	return token, nil
}

func createRandomMessage() (string, error) {
	type CreateRandomMessageResponse struct {
		Data struct {
			CreateRandomMessage string `json:"createRandomMessage"`
		} `json:"data"`
	}

	respBytes, err := utils.CallPostHttpApi(AXIE_INFINITY_GRAPHQL_GATEWAY, GraphqlRequest{
		OperationName: "CreateRandomMessage",
		Query:         "mutation CreateRandomMessage{createRandomMessage}",
	}, nil)

	if err != nil {
		return "", err
	}

	//log.Println(string(respBytes))

	var randomMsgResp CreateRandomMessageResponse
	if err := json.Unmarshal(respBytes, &randomMsgResp); err != nil {
		return "", err
	}

	if randomMsgResp.Data.CreateRandomMessage == "" {
		return "", errors.New("create random message response was empty")
	}

	//log.Printf("Created random message successfully %s", randomMsgResp.Data.CreateRandomMessage)

	return randomMsgResp.Data.CreateRandomMessage, nil
}
