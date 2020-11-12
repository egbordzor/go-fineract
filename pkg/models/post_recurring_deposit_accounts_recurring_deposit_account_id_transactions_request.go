package models

// PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest
type PostRecurringDepositAccountsRecurringDepositAccountIdTransactionsRequest struct {
	Locale            string  `json:"locale,omitempty"`
	DateFormat        string  `json:"dateFormat,omitempty"`
	TransactionDate   string  `json:"transactionDate,omitempty"`
	TransactionAmount float64 `json:"transactionAmount,omitempty"`
	PaymentTypeId     int32   `json:"paymentTypeId,omitempty"`
	AccountNumber     string  `json:"accountNumber,omitempty"`
	CheckNumber       string  `json:"checkNumber,omitempty"`
	RoutingCode       string  `json:"routingCode,omitempty"`
	ReceiptNumber     string  `json:"receiptNumber,omitempty"`
	BankNumber        string  `json:"bankNumber,omitempty"`
}
