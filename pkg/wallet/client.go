package wallet

import (
	"context"
	"fmt"
	"strings"

	"github.com/earn-alliance/wallet-commander-cli/pkg/abi"
	"github.com/earn-alliance/wallet-commander-cli/pkg/constants"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

type Client struct {
	ethClient *ethclient.Client
}

type Balance struct {
	SLP  uint64
	WETH uint64
	AXS  uint64
	RON  uint64
}

func New() (*Client, error) {
	rpcClient, err := rpc.DialHTTP(constants.RONIN_PROVIDER_RPC_URI)

	if err != nil {
		return nil, err
	}

	rpcClient.SetHeader("context-type", "application/json")
	rpcClient.SetHeader("user-agent", "wallet-commander")
	c := ethclient.NewClient(rpcClient)

	return &Client{ethClient: c}, nil
}

func getAddress(address string) common.Address {
	normalizedAddr := address
	if strings.HasPrefix(address, "ronin:") {
		normalizedAddr = strings.ReplaceAll(address, "ronin:", "0x")
	} else if !strings.HasPrefix(address, "0x") {
		normalizedAddr = fmt.Sprintf("0x%s", address)
	}
	return common.HexToAddress(normalizedAddr)
}

func (c *Client) getBalance(contract, addr string) (uint64, error) {
	walletAddr := getAddress(addr)
	axsAddress := common.HexToAddress(contract)

	instance, err := abi.NewRoninBalance(axsAddress, c.ethClient)
	if err != nil {
		return 0, errors.Wrap(err, "unable to initiate RoninBalance contract")
	}

	balance, err := instance.BalanceOf(&bind.CallOpts{}, walletAddr)
	if err != nil {
		return 0, errors.Wrap(err, "unable to call BalanceOf rpc")
	}

	return balance.Uint64(), nil
}

func (c *Client) GetRON(addr string) (uint64, error) {
	address := getAddress(addr)
	balance, err := c.ethClient.BalanceAt(context.Background(), address, nil)
	return balance.Uint64(), err
}

func (c *Client) GetSLP(addr string) (uint64, error) {
	return c.getBalance(constants.SLP_CONTRACT, addr)
}

func (c *Client) GetAXS(addr string) (uint64, error) {
	return c.getBalance(constants.AXS_CONTRACT, addr)
}

func (c *Client) GetWETH(addr string) (uint64, error) {
	return c.getBalance(constants.WETH_CONTRACT, addr)
}

func (c *Client) GetAllBalance(addr string) (Balance, error) {
	slp, err := c.GetSLP(addr)
	if err != nil {
		return Balance{}, errors.Wrap(err, "unable to get SLP contract")
	}

	axs, err := c.GetAXS(addr)
	if err != nil {
		return Balance{}, errors.Wrap(err, "unable to get AXS contract")
	}

	weth, err := c.GetWETH(addr)
	if err != nil {
		return Balance{}, errors.Wrap(err, "unable to get WETH contract")
	}

	ron, err := c.GetRON(addr)
	if err != nil {
		return Balance{}, errors.Wrap(err, "unable to get RON")
	}

	return Balance{
		SLP:  slp,
		WETH: weth,
		AXS:  axs,
		RON:  ron,
	}, nil
}
