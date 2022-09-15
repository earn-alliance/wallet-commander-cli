package api

import validation "github.com/go-ozzo/ozzo-validation/v4"

type WalletCommanderOperation string

const (
	OperationTransferSLP       WalletCommanderOperation = "TRANSFER_SLP"
	OperationBreedAxie         WalletCommanderOperation = "BREED_AXIE"
	OperationTransferAxie      WalletCommanderOperation = "TRANSFER_AXIE"
	OperationAuctionAxie       WalletCommanderOperation = "AUCTION_AXIE"
	OperationCancelAxieAuction WalletCommanderOperation = "CANCEL_AXIE_AUCTION"
	OperationHatchAxie         WalletCommanderOperation = "HATCH_AXIE"
	OperationClaimSLP          WalletCommanderOperation = "CLAIM_SLP"
	OperationCreateAxieAccount WalletCommanderOperation = "CREATE_AXIE_ACCOUNT"
)

type WalletCommanderStatus string

const (
	// When something is successful
	CommandStatusSuccess WalletCommanderStatus = "SUCCESS"
	// When an error has occured, usually coupled with a message
	CommandStatusError WalletCommanderStatus = "ERROR"
	// When the wallet commander has an invalid config and the command cannot run due to lack of permissions or info
	// An example may be that a private key does not exist for the target operation
	// Another example may be that the target account to transferred SLP is not authorized
	CommandStatusBadConfig WalletCommanderStatus = "BAD_CONFIG"
	// For some reason the payload is corrupt or not compatible with this wallet commander
	CommandStatusInvalidPayload WalletCommanderStatus = "INVALID_PAYLOAD"
	// When validating that a transaction was successful, it may timeout after some time because it did not show up on the blockchain
	CommandStatusTimeout WalletCommanderStatus = "TIMEOUT"
	// Default state for a command's status
	CommandStatusNotProcessed WalletCommanderStatus = "NOT_PROCESSED"
	// Status once a command has been locked
	CommandStatusProcessing WalletCommanderStatus = "PROCESSING"
)

type ClaimSlpPayload struct {
	AddressToClaim string `json:"address_to_claim"`
}

type CreateAxieAccountPayload struct {
	Count int `json:"count"`
}

type TransferSlpType string

const (
	TransferSlpTypeScholar TransferSlpType = "SCHOLAR"
	TransferSlpTypeManager TransferSlpType = "MANAGER"
	TransferSlpTypeTrainer TransferSlpType = "TRAINER"
)

type TransferPayload struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type TransferSlpPayload struct {
	TransferPayload
	Amount int             `json:"amount"`
	Type   TransferSlpType `json:"type"`
}

type TransferAxiePayload struct {
	TransferPayload
	AxieId int `json:"axie_id"`
}

func (t TransferPayload) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.To, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&t.From, validation.Required, validation.NilOrNotEmpty))
}

func (t TransferAxiePayload) Validate() error {
	if err := t.TransferPayload.Validate(); err != nil {
		return err
	}

	return validation.ValidateStruct(&t,
		validation.Field(&t.AxieId, validation.Required, validation.Min(1)),
	)
}
