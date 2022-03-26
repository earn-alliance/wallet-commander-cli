package server

import (
	"encoding/json"
	"github.com/earn-alliance/wallet-commander-cli/internal/controller"
	"github.com/earn-alliance/wallet-commander-cli/internal/log"
	"github.com/earn-alliance/wallet-commander-cli/internal/query"
	"github.com/hasura/go-graphql-client"
	"sync"
	"time"
)

type QueryResponse struct {
	WalletCommanderCommands []query.WalletCommanderCommand `json:"wallet_commander_commands"`
}

type Server interface {
	Start()
}

type WalletCommanderServer struct {
	controller controller.Controller
	client     *graphql.SubscriptionClient
	// Temp hack to resolve in future
	commandProcessingLock sync.Map
	graphqlServerEndpoint string
	clientId              string
}

func New(controller controller.Controller, graphqlServerEndpoint, clientId string) (*WalletCommanderServer, error) {
	client := graphql.NewSubscriptionClient(graphqlServerEndpoint).
		WithRetryTimeout(time.Hour).
		WithLog(func(args ...interface{}) {
			log.Logger().Debugf("[GRAPHQL CLIENT] %v", args)
		})

	client.OnError(func(sc *graphql.SubscriptionClient, err error) error {
		log.Logger().Errorf("Graphql client error occurred %v", err)
		return err
	})

	client.OnConnected(func() {
		log.Logger().Infof("Connection to wallet commander server established!")
	})

	client.OnDisconnected(func() {
		log.Logger().Warnf("Wallet commander server has disconnected...")
	})

	return &WalletCommanderServer{
		controller:            controller,
		client:                client,
		graphqlServerEndpoint: graphqlServerEndpoint,
		clientId:              clientId,
	}, nil
}

func (w *WalletCommanderServer) Start() {
	defer w.client.Close()

	err := w.registerWalletCommanderCommandsSubscription()

	if err != nil {
		log.Logger().Panicf("Error creating wallet commander subscriotion %v", err)
	}

	if err := w.client.Run(); err != nil {
		log.Logger().Errorf("Error running graphql subscription client %v", err)
	}
}

func (w *WalletCommanderServer) registerWalletCommanderCommandsSubscription() error {
	_, err := w.client.Subscribe(
		&query.WallterCommanderSubscription,
		query.NewWalletCommanderSubscriptionVars(w.clientId),
		func(message *json.RawMessage, err error) error {
			if err != nil {
				log.Logger().Errorf("Error occurred %v", err)
				return err
			}

			var queryResponse QueryResponse
			if err := json.Unmarshal(*message, &queryResponse); err != nil {
				log.Logger().Errorf("Could not unmarshal graphql message with err %v", err)
				return err
			}

			// Limit requests to 5 per second to avoid spam
			for _, command := range queryResponse.WalletCommanderCommands {
				localCmd := command

				// TODO: Hack that prevents multiple commands from running. This should be done in a custom processosr queue
				_, loaded := w.commandProcessingLock.LoadOrStore(localCmd.Id, true)

				// Already has been processed, so don't do anything
				if loaded {
					continue
				}

				w.controller.ProcessWalletCommand(localCmd)
				// TODO: Hack - this should be done in a custom processor queue
				time.Sleep(time.Millisecond * 1000)
			}

			return nil
		},
	)

	if err == nil {
		log.Logger().Infof("Subscribing to wallet commander server %s with clientId %s", w.graphqlServerEndpoint, w.clientId)
	}

	return err
}
