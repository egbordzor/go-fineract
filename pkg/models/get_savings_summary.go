package models

// GetSavingsSummary struct for GetSavingsSummary
type GetSavingsSummary struct {
	Currency       GetSavingsCurrency `json:"currency,omitempty"`
	AccountBalance int32              `json:"accountBalance,omitempty"`
}
