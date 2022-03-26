package cli

import (
	"github.com/earn-alliance/wallet-commander-cli/internal/controller"
	"github.com/earn-alliance/wallet-commander-cli/internal/log"
	"github.com/earn-alliance/wallet-commander-cli/internal/server"
	"github.com/earn-alliance/wallet-commander-cli/internal/store"
	"github.com/earn-alliance/wallet-commander-cli/internal/types"
	"github.com/earn-alliance/wallet-commander-cli/internal/vault"
	"github.com/earn-alliance/wallet-commander-cli/pkg/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	)

	rootCmd.PersistentFlags().String("log-level", "info", "Set log level to display specific CLI information"+
		"Valid values(trace,debug,info,warn,warning,error,fatal,panic)")

	err := rootCmd.Execute()

	if err != nil {
		log.Logger().Errorf("Error occurred executing CLI command %v", err)
	}
}

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

			vault, err := vault.New(startFlags.SecretsFile)

			if err != nil {
				log.Logger().Panicf("Error occured starting earn alliance commander %v", err)
			}

			client, err := client.New()

			if err != nil {
				log.Logger().Panicf("Error occured starting earn alliance commander %v", err)
			}

			var c controller.Controller

			if startFlags.DevMode == "" {
				c, err = controller.New(vault, store.New(startFlags.EarnAllianceEndpoint), client)
				if err != nil {
					log.Logger().Panicf("Could not start earn alliance commander with err %v", err)
				}
			} else {
				c = controller.NewDevController(startFlags.EarnAllianceEndpoint, startFlags.DevMode)
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
