package models

// GetFixedDepositAccountsAccountIdSummary struct for GetFixedDepositAccountsAccountIdSummary
type GetFixedDepositAccountsAccountIdSummary struct {
	Currency       GetFixedDepositAccountsAccountIdCurrency `json:"currency,omitempty"`
	AccountBalance float32                                  `json:"accountBalance,omitempty"`
}
