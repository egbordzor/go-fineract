package models

// GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTemplateResponse GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTemplateResponse
type GetRecurringDepositAccountsRecurringDepositAccountIdTransactionsTemplateResponse struct {
	Id                 int32                       `json:"id,omitempty"`
	TransactionType    GetRecurringTransactionType `json:"transactionType,omitempty"`
	AccountId          int32                       `json:"accountId,omitempty"`
	AccountNo          string                      `json:"accountNo,omitempty"`
	Date               string                      `json:"date,omitempty"`
	Currency           GetRecurringCurrency        `json:"currency,omitempty"`
	Amount             float32                     `json:"amount,omitempty"`
	Reversed           bool                        `json:"reversed,omitempty"`
	PaymentTypeOptions []int32                     `json:"paymentTypeOptions,omitempty"`
}
