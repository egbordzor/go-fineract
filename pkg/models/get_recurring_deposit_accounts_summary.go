package models

// GetRecurringDepositAccountsSummary struct for GetRecurringDepositAccountsSummary
type GetRecurringDepositAccountsSummary struct {
	Currency       GetRecurringDepositAccountsCurrency `json:"currency,omitempty"`
	AccountBalance float32                             `json:"accountBalance,omitempty"`
}
