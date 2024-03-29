package cli

import (
	"context"
	"errors"
	"fmt"
	"github.com/earn-alliance/wallet-commander-cli/internal/controller"
	"github.com/earn-alliance/wallet-commander-cli/internal/log"
	"github.com/earn-alliance/wallet-commander-cli/internal/server"
	"github.com/earn-alliance/wallet-commander-cli/internal/store"
	"github.com/earn-alliance/wallet-commander-cli/internal/types"
	"github.com/earn-alliance/wallet-commander-cli/internal/vault"
	"github.com/earn-alliance/wallet-commander-cli/pkg/client"
	"github.com/earn-alliance/wallet-commander-cli/pkg/earnalliance"
	"github.com/earn-alliance/wallet-commander-cli/pkg/utils"
	"github.com/earn-alliance/wallet-commander-cli/pkg/wallet"
	"github.com/hasura/go-graphql-client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

func Run() {
	var rootCmd = &cobra.Command{
		Use:   "wallet-commander",
		Short: "Execute or proxy blockchain operations for your wallet assets",
		Long:  `Execute or proxy blockchain operations for your wallet assets. Docs at https://earnalliance.com/docs/features/wallet-commander`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	rootCmd.AddCommand(
		startCommand(),
		getAxieCommand(),
	)

	rootCmd.PersistentFlags().String("log-level", "info", "Set log level to display specific CLI information"+
		"Valid values(trace,debug,info,warn,warning,error,fatal,panic)")

	err := rootCmd.Execute()

	if err != nil {
		log.Logger().Errorf("Error occurred executing CLI command %v", err)
	}
}

//func checkNewVersion() {
//	githubTag := &latest.GithubTag{
//		Owner: config.CliGitHubOrg,
//		Repository: config.CliGitHubRepo,
//		FixVersionStrFunc: latest.DeleteFrontV(),
//	}
//
//	res, err := latest.Check(githubTag, config.Version[1:])
//
//	if err != nil {
//		log.Logger().Errorf("Could not check latest version for CLI %v", err)
//	}
//
//	if res.Outdated {
//		log.Logger().Warnf("A new version of the CLI is available. Current version: %s. Latest version: %s",
//			config.Version,
//			res.Current,
//		)
//	}
//}

func startCommand() *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Starts a wallet commander",
		Long:  "Initialize wallet commander to connect to specified 3rd party to proxy and sign blockchain operations",
	}
	// Using sub commands so more platforms can contribute here
	startCmd.AddCommand(
		startEarnAllianceCommander(),
	)

	startCmd.PersistentFlags().String("dev-mode", "", "Enable development mocking mode. "+
		"Valid values (SUCCESS,FAILED_TRANSACTION,INTERNAL_ERROR)")

	startCmd.PersistentFlags().String("earn-alliance-endpoint", types.EarnManagementServerEndpoint, "Earn Management server endpoint")
	startCmd.PersistentFlags().String("secrets-file", types.DefaultSecretsFileName, "Secrets file with array of private keys")

	startCmd.PersistentFlags().String("client-id", "", "Earn Management account Client Id (required)")
	startCmd.MarkPersistentFlagRequired("client-id")

	return startCmd
}

func getAxieCommand() *cobra.Command {
	axieCmd := &cobra.Command{
		Use:   "axie",
		Short: "Runs queries and operations for axie infinity",
		Long:  "Manually query and run operations for axie infinity's blockchain and apis",
	}
	// Using sub commands so more platforms can contribute here
	axieCmd.AddCommand(
		getAxieClaimPayloadCommand(),
		getAxieClaimCommand(),
		createAccountsCommand(),
	)

	return axieCmd
}

func getAxieClaimPayloadCommand() *cobra.Command {
	axieClaimableSlp := &cobra.Command{
		Use:   "claim-payload",
		Short: "Get claim payload payload for a ronin address",
		Long:  "Claim payloads are used for signing against the blockchain for claiming SLP. This command requires secrets.json",
		Run: func(cmd *cobra.Command, args []string) {
			setupLoggerFlags(cmd)
			secretsFile, err := cmd.Flags().GetString("secrets-file")

			if err != nil {
				log.Logger().Panicf("Internal Error: Secrets File flag not setup correctly %v", err)
			}

			client, err := client.New()
			if err != nil {
				log.Logger().Panicf("could not create axie client with err: %v", err)
			}

			vault, err := vault.New(secretsFile)

			if err != nil {
				log.Logger().Fatal("could not create vault %v", err)
			}

			addr, _ := cmd.Flags().GetString("ronin-address")

			privateKey, err := vault.GetPrivateKey(addr)

			if err != nil {
				log.Logger().Panicf("could not get private key with err: %v", err)
			}

			resp, err := client.GetClaimOriginPayload(context.Background(), addr, privateKey)

			if err != nil {
				log.Logger().Panicf("Error occurred getting claimable amount %v", err)
			}

			log.Logger().Printf("Payload authenticated with claimable amount for addr is %v", resp.GetClaimableAmount())
			log.Logger().Printf("Payload json %v", resp)
		},
	}

	axieClaimableSlp.Flags().String("ronin-address", "", "Ronin address to check for claimable slp")
	axieClaimableSlp.MarkFlagRequired("ronin-address")

	axieClaimableSlp.PersistentFlags().String("secrets-file", types.DefaultSecretsFileName, "Secrets file with array of private keys")

	return axieClaimableSlp
}

func getAxieClaimCommand() *cobra.Command {
	axieClaimableSlp := &cobra.Command{
		Use:   "claim-slp",
		Short: "Claim's SLP for a ronin address",
		Long:  "Runs claiming logic for ronin address. This command requires secrets.json",
		Run: func(cmd *cobra.Command, args []string) {
			setupLoggerFlags(cmd)
			secretsFile, err := cmd.Flags().GetString("secrets-file")

			if err != nil {
				log.Logger().Panicf("Internal Error: Secrets File flag not setup correctly %v", err)
			}

			client, err := client.New()
			if err != nil {
				log.Logger().Panicf("could not create axie client with err: %v", err)
			}

			vault, err := vault.New(secretsFile)

			if err != nil {
				log.Logger().Fatal("could not create vault %v", err)
			}

			addr, _ := cmd.Flags().GetString("ronin-address")

			privateKey, err := vault.GetPrivateKey(addr)

			if err != nil {
				log.Logger().Panicf("could not get private key with err: %v", err)
			}

			resp, err := client.ClaimOriginSlp(context.Background(), utils.RoninAddrToEthAddr(addr), privateKey)

			if err != nil {
				log.Logger().Panicf("Error occurred getting claimable amount %v", err)
			}

			log.Logger().Printf("Claim tx %v", resp)
		},
	}

	axieClaimableSlp.Flags().String("ronin-address", "", "Ronin address to check for claimable slp")
	axieClaimableSlp.MarkFlagRequired("ronin-address")

	axieClaimableSlp.PersistentFlags().String("secrets-file", types.DefaultSecretsFileName, "Secrets file with array of private keys")

	return axieClaimableSlp
}

func createAccountsCommand() *cobra.Command {
	command := &cobra.Command{
		Long:  "Create Axie Infinity Login Accounts for a client. This command requires client-id",
		Use:   "create-accounts {number-of-accounts}",
		Short: "Create verified Axie Infinity login accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			setupLoggerFlags(cmd)

			clientId, err := cmd.Flags().GetString("client-id")
			secretsFile, err := cmd.Flags().GetString("secrets-file")

			if err != nil {
				//log.Logger().Errorf("Error: client-id flag not setup correctly %v", err)
				return errors.New("client-id flag not setup correctly")
			}

			if len(args) < 1 {
				return errors.New("incorrect number of arguments")
			}

			count, err := strconv.Atoi(args[0])

			if err != nil || count <= 0 {
				return errors.New(fmt.Sprintf("Error: Incorrect number of accounts value: %s", args[0]))
			}

			graphqlServerEndpoint, err := cmd.Flags().GetString("earn-alliance-endpoint")

			clientVault, err := vault.New(secretsFile)

			if err != nil {
				log.Logger().Panicf("Error occured starting earn alliance commander %v", err)
			}

			axieClient, err := client.New()

			if err != nil {
				log.Logger().Panicf("Error occured starting earn alliance commander %v", err)
			}

			walletClient, err := wallet.New()
			if err != nil {
				log.Logger().Panicf("Error occured starting earn alliance commander %v", err)
			}

			gqlClient := graphql.NewClient(graphqlServerEndpoint, nil)
			eaClient := earnalliance.New(clientId, axieClient, walletClient, gqlClient)

			return eaClient.CreateAxieInfinityAccount(count, clientVault)
		},
	}

	command.PersistentFlags().String("earn-alliance-endpoint", types.EarnManagementServerEndpoint, "Earn Management server endpoint")
	command.PersistentFlags().String("client-id", "", "Earn Management account Client Id (required)")
	command.PersistentFlags().String("secrets-file", types.DefaultSecretsFileName, "Secrets file with array of private keys")
	command.MarkPersistentFlagRequired("client-id")

	return command
}

type StartFlags struct {
	EarnAllianceEndpoint string
	SecretsFile          string
	DevMode              types.DevMockMode
	ClientId             string
}

func setupLoggerFlags(cmd *cobra.Command) {
	levelStr, err := cmd.Flags().GetString("log-level")

	if err != nil {
		log.Logger().Panicf("Internal Error: Log Level flag not setup correctly")
	}

	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		log.Logger().Panicf("Incorrect log-level value provided! %v", err)
	}

	log.Logger().SetLevel(level)
}

func getStartFlags(cmd *cobra.Command) StartFlags {
	secretsFile, err := cmd.Flags().GetString("secrets-file")

	if err != nil {
		log.Logger().Panicf("Internal Error: Secrets File flag not setup correctly %v", err)
	}

	devModeStr, err := cmd.Flags().GetString("dev-mode")

	if err != nil {
		log.Logger().Panicf("Internal Error: Dev Mode flag not setup correctly %v", err)
	}

	earnAllianceEndpoint, err := cmd.Flags().GetString("earn-alliance-endpoint")

	if err != nil {
		log.Logger().Panicf("Internal Error: Earn Alliance Endpoint flag not setup correctly %v", err)
	}

	clientId, err := cmd.Flags().GetString("client-id")

	if err != nil {
		log.Logger().Panicf("Internal Error: Client Id flag not setup correctly %v", err)
	}

	return StartFlags{
		EarnAllianceEndpoint: earnAllianceEndpoint,
		SecretsFile:          secretsFile,
		DevMode:              types.DevMockMode(devModeStr),
		ClientId:             clientId,
	}
}

func startEarnAllianceCommander() *cobra.Command {
	var startEarnManagementCommander = &cobra.Command{
		Use:   "earn-alliance --clientId {id}",
		Short: "Starts earn alliance wallet commander",
		Long:  `echo things multiple times back to the user by providing a count and a string.`,
		Run: func(cmd *cobra.Command, args []string) {
			setupLoggerFlags(cmd)
			startFlags := getStartFlags(cmd)

			err, c := createController(cmd, startFlags)

			if err != nil {
				log.Logger().Panicf("Could not create controller %v", err)
			}

			server, err := server.New(c, startFlags.EarnAllianceEndpoint, startFlags.ClientId)

			if err != nil {
				log.Logger().Panicf("Could not start earn alliance commander with err %v", err)
			}

			server.Start()
		},
	}

	return startEarnManagementCommander
}

func createController(cmd *cobra.Command, startFlags StartFlags) (error, controller.Controller) {

	vault, err := vault.New(startFlags.SecretsFile)

	if err != nil {
		log.Logger().Panicf("Error occured starting earn alliance commander %v", err)
	}

	client, err := client.New()

	if err != nil {
		log.Logger().Panicf("Error occured starting earn alliance commander %v", err)
	}

	walletClient, err := wallet.New()
	if err != nil {
		log.Logger().Panicf("Error occured starting earn alliance commander %v", err)
	}

	gqlClient := graphql.NewClient(startFlags.EarnAllianceEndpoint, nil)

	eaClient := earnalliance.New(startFlags.ClientId, client, walletClient, gqlClient)
	var c controller.Controller

	if startFlags.DevMode == "" {
		c, err = controller.New(vault, store.New(gqlClient), client, eaClient)
		if err != nil {
			log.Logger().Panicf("Could not start earn alliance commander with err %v", err)
		}
	} else {
		c = controller.NewDevController(store.New(gqlClient), startFlags.DevMode)
	}
	return err, c
}
