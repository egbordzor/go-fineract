package models

// GetSelfSavingsSummary struct for GetSelfSavingsSummary
type GetSelfSavingsSummary struct {
	Currency       GetSelfSavingsCurrency `json:"currency,omitempty"`
	AccountBalance int32                  `json:"accountBalance,omitempty"`
}
