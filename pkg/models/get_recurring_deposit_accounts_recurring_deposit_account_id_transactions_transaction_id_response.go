package models

// GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTransactionIdResponse GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTransactionIdResponse
type GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTransactionIdResponse struct {
	Id                int32                                   `json:"id,omitempty"`
	TransactionType   GetRecurringTransactionsTransactionType `json:"transactionType,omitempty"`
	AccountId         int32                                   `json:"accountId,omitempty"`
	AccountNo         string                                  `json:"accountNo,omitempty"`
	Date              string                                  `json:"date,omitempty"`
	Currency          GetRecurringTransactionsCurrency        `json:"currency,omitempty"`
	PaymentDetailData GetRecurringPaymentDetailData           `json:"paymentDetailData,omitempty"`
	Amount            float32                                 `json:"amount,omitempty"`
	RunningBalance    int32                                   `json:"runningBalance,omitempty"`
	Reversed          bool                                    `json:"reversed,omitempty"`
}
