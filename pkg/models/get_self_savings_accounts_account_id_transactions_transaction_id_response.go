package models

// GetSelfSavingsAccountsAccountIdTransactionsTransactionIdResponse GetSelfSavingsAccountsAccountIdTransactionsTransactionIdResponse
type GetSelfSavingsAccountsAccountIdTransactionsTransactionIdResponse struct {
	Id                int32                             `json:"id,omitempty"`
	TransactionType   GetSelfSavingsTransactionType     `json:"transactionType,omitempty"`
	AccountId         int32                             `json:"accountId,omitempty"`
	AccountNo         int64                             `json:"accountNo,omitempty"`
	Date              string                            `json:"date,omitempty"`
	Currency          GetSelfSavingsTransactionCurrency `json:"currency,omitempty"`
	PaymentDetailData GetSelfSavingsPaymentDetailData   `json:"paymentDetailData,omitempty"`
	Amount            int32                             `json:"amount,omitempty"`
	RunningBalance    int32                             `json:"runningBalance,omitempty"`
	Reversed          bool                              `json:"reversed,omitempty"`
}
