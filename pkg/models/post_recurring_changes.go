package models

// PostRecurringChanges struct for PostRecurringChanges
type PostRecurringChanges struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	CheckNumber   string `json:"checkNumber,omitempty"`
	RoutingCode   string `json:"routingCode,omitempty"`
	ReceiptNumber string `json:"receiptNumber,omitempty"`
	BankNumber    string `json:"bankNumber,omitempty"`
}
