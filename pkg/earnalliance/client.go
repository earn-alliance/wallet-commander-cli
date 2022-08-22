package earnalliance

import (
	"context"
	"errors"
	"fmt"
	"github.com/earn-alliance/wallet-commander-cli/internal/log"
	"github.com/earn-alliance/wallet-commander-cli/internal/query"
	"github.com/earn-alliance/wallet-commander-cli/internal/vault"
	"github.com/earn-alliance/wallet-commander-cli/pkg/client"
	"github.com/earn-alliance/wallet-commander-cli/pkg/utils"
	"github.com/earn-alliance/wallet-commander-cli/pkg/wallet"
	"github.com/hasura/go-graphql-client"
)

type EarnAllianceClient struct {
	gqlClient    *graphql.Client
	axieClient   client.Client
	walletClient *wallet.Client
	clientId     string
}

func New(graphqlServerEndpoint, clientId string) (*EarnAllianceClient, error) {
	axieClient, err := client.New()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Internal Error: could not create axie client with err %v", err))
	}

	walletClient, err := wallet.New()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Internal Error: could not create wallet client with err %v", err))
	}

	return &EarnAllianceClient{
		gqlClient:    graphql.NewClient(graphqlServerEndpoint, nil),
		clientId:     clientId,
		axieClient:   axieClient,
		walletClient: walletClient,
	}, nil
}

func (c *EarnAllianceClient) CreateAxieInfinityAccount(count int, vault vault.Vault) error {

	var batchSize = 5
	var accounts []query.AxieInfinityAccount

	for i := 0; i < count; i++ {
		newWallet, err := c.walletClient.CreateWallet()
		if err != nil {
			return errors.New(fmt.Sprintf("Internal Error: could not create wallet with err %v", err))
		}

		accessToken, err := c.axieClient.CreateAccessToken(newWallet.Address, newWallet.PrivateKey)
		if err != nil {
			return errors.New(fmt.Sprintf("Internal Error: could not access token with err %v", err))
		}

		accounts = append(accounts, query.AxieInfinityAccount{
			RoninWallet: utils.EthAddrToRoninAddr(newWallet.Address),
			AccessToken: accessToken,
		})

		vault.AddRoninWallet(utils.EthAddrToRoninAddr(newWallet.Address), newWallet.PrivateKey)

		if i > 0 && i%batchSize == 0 {
			log.Logger().Infof("[QUEUE] Enqueuing %d of %d requests", i, count)

			err := c.insertAxieInfinityAccounts(accounts)
			if err != nil {
				return errors.New(fmt.Sprintf("Internal Error: could not enqueue request %v", err))
			}

			if err := vault.Save(); err != nil {
				return errors.New(fmt.Sprintf("Internal Error: unable to update secret file %v", err))
			}
			accounts = []query.AxieInfinityAccount{}
		}
	}

	if len(accounts) > 0 {
		err := c.insertAxieInfinityAccounts(accounts)
		if err != nil {
			return errors.New(fmt.Sprintf("Internal Error: could not enqueue request %v", err))
		}

		if err := vault.Save(); err != nil {
			return errors.New(fmt.Sprintf("Internal Error: unable to update secret file %v", err))
		}
	}

	log.Logger().Infof("[QUEUE] Successfully enqueued %d Axie Infinity account creation request", count)
	log.Logger().Infof("[VAULT] Successfully updated secrets file. %d wallets in total", vault.GetWalletCounts())
	log.Logger().Infof("[VAULT] Make sure to backup this file with your new private keys!")
	return nil
}

func (c *EarnAllianceClient) insertAxieInfinityAccounts(accounts []query.AxieInfinityAccount) error {
	err := c.gqlClient.Mutate(context.Background(), &query.InsertAxieInfinityAccounts, query.NewInsertAxieInfinityAccountsVars(
		c.clientId,
		accounts,
	))

	if err != nil {
		return errors.New(fmt.Sprintf("could not create new axie infinity account with err: %v", err))
	}
	return nil
}
