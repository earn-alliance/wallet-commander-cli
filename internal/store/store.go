package store

import (
	"context"
	"github.com/earn-alliance/wallet-commander-cli/internal/log"
	"github.com/earn-alliance/wallet-commander-cli/internal/query"
	"github.com/earn-alliance/wallet-commander-cli/pkg/api"
	"github.com/hasura/go-graphql-client"
)

type Store interface {
	UpdateWalletCommandTransactionSuccess(id string, transaction string)
	UpdateWalletCommandTransactionError(id string, status api.WalletCommanderStatus, msg string)
}

type WalletCommanderStore struct {
	client *graphql.Client
}

func (w *WalletCommanderStore) UpdateWalletCommandTransactionSuccess(id string, transaction string) {
	err := w.client.Mutate(context.Background(), &query.UpdateCommandResults, query.NewUpdateCommandResultsVars(
		id,
		transaction,
		api.CommandStatusSuccess,
		"",
	))

	if err != nil {
		log.Logger().Errorf("could not update wallet command transaction to a successful state with err: %v", err)
	}
}

func (w *WalletCommanderStore) UpdateWalletCommandTransactionError(id string, status api.WalletCommanderStatus, msg string) {
	err := w.client.Mutate(context.Background(), &query.UpdateCommandResults, query.NewUpdateCommandResultsVars(
		id,
		"",
		status,
		msg,
	))

	if err != nil {
		log.Logger().Errorf("could not update wallet command transaction to a errored state with err: %v", err)
	}
}

func New(graphqlServerEndpoint string) Store {
	return &WalletCommanderStore{
		client: graphql.NewClient(graphqlServerEndpoint, nil),
	}
}
