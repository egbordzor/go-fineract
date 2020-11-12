package models

// GetFixedDepositAccountsSummary struct for GetFixedDepositAccountsSummary
type GetFixedDepositAccountsSummary struct {
	Currency       GetFixedDepositAccountsCurrency `json:"currency,omitempty"`
	AccountBalance float32                         `json:"accountBalance,omitempty"`
}
