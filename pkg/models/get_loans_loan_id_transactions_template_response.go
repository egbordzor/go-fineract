package models

// GetLoansLoanIdTransactionsTemplateResponse GetLoansLoanIdTransactionsTemplateResponse
type GetLoansLoanIdTransactionsTemplateResponse struct {
	TransactionType GetLoansTransactionType `json:"transactionType,omitempty"`
	Date            string                  `json:"date,omitempty"`
	Total           GetLoansTotal           `json:"total,omitempty"`
}
