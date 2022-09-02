package query

import (
	"encoding/json"
	"errors"
	"github.com/earn-alliance/wallet-commander-cli/pkg/api"
	"github.com/hasura/go-graphql-client"
	"time"
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

var UpsertWalletCommanderActiveClients struct {
	InsertWalletCommanderActiveClients struct {
		Client_Id string
	} `graphql:"insert_wallet_commander_active_clients_one(object:{ client_id: $clientId, last_ping_time: $lastPingTime, last_connected_status: $connected }, on_conflict:{ constraint:wallet_commander_active_clients_pkey, update_columns:[client_id,last_ping_time,last_connected_status]})"`
}

func NewUpsertWalletCommanderActiveClientVars(clientId string, connected bool) map[string]interface{} {
	return map[string]interface{}{
		"clientId":     uuid(clientId),
		"connected":    graphql.Boolean(connected),
		"lastPingTime": timestamptz(time.Now().Format("2006-01-02 15:04:05")),
	}
}

type AxieInfinityAccount struct {
	RoninWallet string
	AccessToken string
}

var InsertAxieInfinityAccounts struct {
	InsertAxieInfinityAccounts struct {
		Affected_Rows int
	} `graphql:"insert_axie_infinity_accounts(objects: $objects)"`
}

func NewInsertAxieInfinityAccountsVars(clientId string, accounts []AxieInfinityAccount) map[string]interface{} {
	var objects []axie_infinity_accounts_insert_input

	for _, acc := range accounts {
		objects = append(objects, axie_infinity_accounts_insert_input{
			"client_id":    uuid(clientId),
			"ronin_wallet": graphql.String(acc.RoninWallet),
			"access_token": graphql.String(acc.AccessToken),
		})
	}

	return map[string]interface{}{
		"objects": objects,
	}
}

type WalletCommanderCommand struct {
	Id        graphql.ID
	Operation graphql.String
	Payload   json.RawMessage
}

func (w WalletCommanderCommand) UnmarshalCreateAxieAccountPayload() (*api.CreateAxieAccountPayload, error) {
	if string(w.Operation) != string(api.OperationCreateAxieAccount) {
		return nil, errors.New("cannot unmarshal CreateAxieAccountPayload because operation type is not CREATE_AXIE_ACCOUNT")
	}

	var createAxieAccountPayload api.CreateAxieAccountPayload
	if err := json.Unmarshal(w.Payload, &createAxieAccountPayload); err != nil {
		return nil, err
	}

	return &createAxieAccountPayload, nil
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
