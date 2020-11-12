package models

// PostCollectionSheetRequest PostCollectionSheetRequest
type PostCollectionSheetRequest struct {
	DateFormat                   string                                       `json:"dateFormat,omitempty"`
	Locale                       string                                       `json:"locale,omitempty"`
	TransactionDate              string                                       `json:"transactionDate,omitempty"`
	ActualDisbursementDate       string                                       `json:"actualDisbursementDate,omitempty"`
	BulkDisbursementTransactions []int32                                      `json:"bulkDisbursementTransactions,omitempty"`
	BulkRepaymentTransactions    PostCollectionSheetBulkRepaymentTransactions `json:"bulkRepaymentTransactions,omitempty"`
	BulkSavingsDueTransactions   []int32                                      `json:"bulkSavingsDueTransactions,omitempty"`
}
