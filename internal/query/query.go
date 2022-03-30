package query

import (
	"encoding/json"
	"errors"
	"github.com/earn-alliance/wallet-commander-cli/pkg/api"
	"github.com/hasura/go-graphql-client"
)

var WallterCommanderSubscription struct {
	WalletCommanderCommand `graphql:"wallet_commander_commands(where: {client_id: {_eq: $clientId }, command_status:{ _eq: NOT_PROCESSED }})"`
}

func NewWalletCommanderSubscriptionVars(clientId string) map[string]interface{} {
	return map[string]interface{}{
		"clientId": uuid(clientId),
	}
}

var UpdateCommandResults struct {
	UpdateWalletCommanderCommands struct {
		Id uuid
	} `graphql:"update_wallet_commander_commands_by_pk(pk_columns: {id: $id}, _set: {completed_transaction_hash: $transactionHash, command_status: $status, command_status_message: $message})"`
}

func NewUpdateCommandResultsVars(
	commandId, transactionHash string, status api.WalletCommanderStatus, statusMessage string) map[string]interface{} {
	return map[string]interface{}{
		"id":              uuid(commandId),
		"transactionHash": graphql.String(transactionHash),
		"status":          wallet_commander_command_statuses_enum(status),
		"message":         graphql.String(statusMessage),
	}
}

type WalletCommanderCommand struct {
	Id        graphql.ID
	Operation graphql.String
	Payload   json.RawMessage
}

func (w WalletCommanderCommand) UnmarshalClaimSlpPayload() (*api.ClaimSlpPayload, error) {
	if string(w.Operation) != string(api.OperationClaimSLP) {
		return nil, errors.New("cannot unmarshal ClaimSlpPayload because operation type is not CLAIM_SLP")
	}

	var claimSlpPayload api.ClaimSlpPayload
	if err := json.Unmarshal(w.Payload, &claimSlpPayload); err != nil {
		return nil, err
	}

	return &claimSlpPayload, nil
}

func (w WalletCommanderCommand) UnmarshalTransferSlpPayload() (*api.TransferSlpPayload, error) {
	if string(w.Operation) != string(api.OperationTransferSLP) {
		return nil, errors.New("cannot unmarshal TransferSlpPayload because operation type is not TRANSFER_SLP")
	}

	var transferSlpPayload api.TransferSlpPayload
	if err := json.Unmarshal(w.Payload, &transferSlpPayload); err != nil {
		return nil, err
	}

	return &transferSlpPayload, nil
}

func (w WalletCommanderCommand) UnmarshalTransferAxiePayload() (*api.TransferAxiePayload, error) {
	if string(w.Operation) != string(api.OperationTransferAxie) {
		return nil, errors.New("cannot unmarshal TransferAxiePayload because operation type is not TRANSFER_AXIE")
	}

	var transferAxiePayload api.TransferAxiePayload
	if err := json.Unmarshal(w.Payload, &transferAxiePayload); err != nil {
		return nil, err
	}

	return &transferAxiePayload, nil
}
