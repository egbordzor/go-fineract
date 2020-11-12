package models

// GetLoansLoanIdTransactionsTransactionIdResponse GetLoansLoanIdTransactionsTransactionIdResponse
type GetLoansLoanIdTransactionsTransactionIdResponse struct {
	Id               int32            `json:"id,omitempty"`
	Type             GetLoansType     `json:"type,omitempty"`
	Date             string           `json:"date,omitempty"`
	ManuallyReversed bool             `json:"manuallyReversed,omitempty"`
	Currency         GetLoansCurrency `json:"currency,omitempty"`
	Amount           float64          `json:"amount,omitempty"`
	InterestPortion  float64          `json:"interestPortion,omitempty"`
}
