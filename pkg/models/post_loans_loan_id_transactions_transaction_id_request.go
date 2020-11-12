package models

// PostLoansLoanIdTransactionsTransactionIdRequest PostLoansLoanIdTransactionsTransactionIdRequest
type PostLoansLoanIdTransactionsTransactionIdRequest struct {
	Locale            string  `json:"locale,omitempty"`
	DateFormat        string  `json:"dateFormat,omitempty"`
	TransactionDate   string  `json:"transactionDate,omitempty"`
	TransactionAmount float64 `json:"transactionAmount,omitempty"`
	Note              string  `json:"note,omitempty"`
}
