package models

// GetSelfLoansLoanIdTransactionsTransactionIdResponse GetSelfLoansLoanIdTransactionsTransactionIdResponse
type GetSelfLoansLoanIdTransactionsTransactionIdResponse struct {
	Id               int32                              `json:"id,omitempty"`
	Type             GetSelfLoansLoanIdTransactionsType `json:"type,omitempty"`
	Date             string                             `json:"date,omitempty"`
	ManuallyReversed bool                               `json:"manuallyReversed,omitempty"`
	Currency         GetLoanCurrency                    `json:"currency,omitempty"`
	Amount           float32                            `json:"amount,omitempty"`
	InterestPortion  float32                            `json:"interestPortion,omitempty"`
}
