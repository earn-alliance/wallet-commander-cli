package types

type WalletTransactionHistory struct {
	Total   int `json:"total"`
	Results []struct {
		From          string `json:"from" csv:"from"`
		To            string `json:"to" csv:"to"`
		Value         string `json:"value" csv:"value"`
		Timestamp     int64  `json:"timestamp" csv:"timestamp"`
		TimestampStr  string `json:"timestamp_string" csv:"timestamp_string"`
		LogIndex      string `json:"log_index"`
		TxHash        string `json:"tx_hash"`
		TokenAddress  string `json:"token_address"`
		TokenDecimals int    `json:"token_decimals"`
		TokenName     string `json:"token_name"`
		TokenSymbol   string `json:"token_symbol"`
		TokenType     string `json:"token_type"`
	} `json:"results"`
}
