package types

type DevMockMode string

const (
	DevMockModeSuccess           DevMockMode = "SUCCESS"
	DevMockModeFailedTransaction DevMockMode = "FAILED_TRANSACTION"
	DevMockModeInternalError     DevMockMode = "INTERNAL_ERROR"

	EarnManagementServerEndpoint = "https://graphql.earnalliance.com/v1/graphql"
	DefaultSecretsFileName       = "secrets.json"
)
