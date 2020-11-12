package models

// GetSavingsAccountsSummary struct for GetSavingsAccountsSummary
type GetSavingsAccountsSummary struct {
	Currency         GetSavingsCurrency `json:"currency,omitempty"`
	AccountBalance   int32              `json:"accountBalance,omitempty"`
	AvailableBalance int32              `json:"availableBalance,omitempty"`
}
