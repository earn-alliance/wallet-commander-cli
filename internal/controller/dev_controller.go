package controller

import (
	"github.com/earn-alliance/wallet-commander-cli/internal/log"
	"github.com/earn-alliance/wallet-commander-cli/internal/query"
	"github.com/earn-alliance/wallet-commander-cli/internal/store"
	"github.com/earn-alliance/wallet-commander-cli/internal/types"
	"github.com/earn-alliance/wallet-commander-cli/pkg/api"
)

type DevController struct {
	store store.Store
	mode  types.DevMockMode
}

func NewDevController(serverEndpoint string, mode types.DevMockMode) Controller {
	if mode == "" {
		panic("Invalid dev mock mode set in env vars!")
	}

	log.Logger().Infof("Server has been configured in DEV_MODE with the setting of %s", mode)

	return &DevController{
		store: store.New(serverEndpoint),
		mode:  mode,
	}
}

func (d *DevController) ProcessWalletCommand(command query.WalletCommanderCommand) {
	commandId := command.Id.(string)
	log.Logger().Infof("Dev Server in mode %s is processing command %s", d.mode, commandId)
	const failedTransaction = "0xaa3e4b0982118c350b599e8339492aab1472e2b846a73366ef2d6936528d43df"
	const successTransaction = "0x4200af53a36396e9880e328a33c744b0c27178c79a068be58d770df5874a55e5"

	switch d.mode {
	case types.DevMockModeSuccess:
		d.store.UpdateWalletCommandTransactionSuccess(commandId, successTransaction)
		break
	case types.DevMockModeFailedTransaction:
		d.store.UpdateWalletCommandTransactionSuccess(commandId, failedTransaction)
		break
	case types.DevMockModeInternalError:
		d.store.UpdateWalletCommandTransactionError(
			commandId, api.CommandStatusBadConfig, "[Mock] No private key available")
		break
	}
}
