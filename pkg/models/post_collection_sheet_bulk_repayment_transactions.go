package models

// PostCollectionSheetBulkRepaymentTransactions struct for PostCollectionSheetBulkRepaymentTransactions
type PostCollectionSheetBulkRepaymentTransactions struct {
	LoanId            int32   `json:"loanId,omitempty"`
	TransactionAmount float64 `json:"transactionAmount,omitempty"`
	PaymentTypeId     int32   `json:"paymentTypeId,omitempty"`
	ReceiptNumber     int64   `json:"receiptNumber,omitempty"`
}
