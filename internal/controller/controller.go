package controller

import (
	"context"
	"github.com/earn-alliance/wallet-commander-cli/internal/log"
	"github.com/earn-alliance/wallet-commander-cli/internal/query"
	"github.com/earn-alliance/wallet-commander-cli/internal/store"
	"github.com/earn-alliance/wallet-commander-cli/internal/vault"
	"github.com/earn-alliance/wallet-commander-cli/pkg/api"
	"github.com/earn-alliance/wallet-commander-cli/pkg/client"
	"github.com/earn-alliance/wallet-commander-cli/pkg/utils"
)

type Controller interface {
	ProcessWalletCommand(command query.WalletCommanderCommand)
}

type WalletCommanderController struct {
	client client.Client
	vault  vault.Vault
	store  store.Store
}

func New(vault vault.Vault, store store.Store, client client.Client) (Controller, error) {
	return &WalletCommanderController{
		client: client,
		vault:  vault,
		store:  store,
	}, nil
}

func (w *WalletCommanderController) ProcessWalletCommand(command query.WalletCommanderCommand) {
	log.Logger().Infof("Processing command id %s of type %s", command.Id, command.Operation)

	commandIdStr := command.Id.(string)
	switch string(command.Operation) {
	case string(api.OperationClaimSLP):
		if payload, err := command.UnmarshalClaimSlpPayload(); err != nil {
			w.store.UpdateWalletCommandTransactionError(commandIdStr, api.CommandStatusInvalidPayload, err.Error())
		} else {
			w.claimSlp(commandIdStr, payload)
		}
	case string(api.OperationTransferSLP):
		if payload, err := command.UnmarshalTransferSlpPayload(); err != nil {
			w.store.UpdateWalletCommandTransactionError(commandIdStr, api.CommandStatusInvalidPayload, err.Error())
		} else {
			w.transferSlp(commandIdStr, payload)
		}
	case string(api.OperationTransferAxie):
		if payload, err := command.UnmarshalTransferAxiePayload(); err != nil {
			w.store.UpdateWalletCommandTransactionError(commandIdStr, api.CommandStatusInvalidPayload, err.Error())
		} else if err := payload.Validate(); err != nil {
			w.store.UpdateWalletCommandTransactionError(commandIdStr, api.CommandStatusInvalidPayload, err.Error())
		} else {
			w.transferAxie(commandIdStr, payload)
		}
	}
}

func (w *WalletCommanderController) processTransactionResult(commandId string, transaction string, err error) {
	if err == nil {
		log.Logger().Debugf("wrting successful result for command %s", commandId)
		w.store.UpdateWalletCommandTransactionSuccess(commandId, transaction)
	} else {
		log.Logger().Debugf("wrting an error result for command %s with err %v", commandId, err)
		w.store.UpdateWalletCommandTransactionError(commandId, api.CommandStatusError, err.Error())
	}
}

func (w *WalletCommanderController) getPrivateKey(commandId, address string) (string, error) {
	key, err := w.vault.GetPrivateKey(address)

	if err != nil {
		w.store.UpdateWalletCommandTransactionError(commandId, api.CommandStatusBadConfig, err.Error())
	}

	return key, err
}

func (w *WalletCommanderController) claimSlp(commandId string, payload *api.ClaimSlpPayload) {
	key, err := w.vault.GetPrivateKey(payload.AddressToClaim)

	if err != nil {
		w.store.UpdateWalletCommandTransactionError(commandId, api.CommandStatusBadConfig, err.Error())
		return
	}

	tx, err := w.client.ClaimSlp(context.Background(), utils.RoninAddrToEthAddr(payload.AddressToClaim), key)

	w.processTransactionResult(commandId, tx, err)
}

func (w *WalletCommanderController) transferSlp(commandId string, payload *api.TransferSlpPayload) {
	key, err := w.vault.GetPrivateKey(payload.From)

	if err != nil {
		w.store.UpdateWalletCommandTransactionError(commandId, api.CommandStatusBadConfig, err.Error())
		return
	}

	tx, err := w.client.TransferSlp(context.Background(), key, utils.RoninAddrToEthAddr(payload.To), payload.Amount)

	w.processTransactionResult(commandId, tx, err)
}

func (w *WalletCommanderController) transferAxie(commandId string, payload *api.TransferAxiePayload) {
	if key, err := w.getPrivateKey(commandId, payload.From); err == nil {
		tx, err := w.client.TransferAxie(context.Background(), key, utils.RoninAddrToEthAddr(payload.To), payload.AxieId)
		w.processTransactionResult(commandId, tx, err)
	}
}
