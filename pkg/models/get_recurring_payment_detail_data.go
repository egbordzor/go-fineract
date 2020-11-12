package models

// GetRecurringPaymentDetailData struct for GetRecurringPaymentDetailData
type GetRecurringPaymentDetailData struct {
	Id            int32                   `json:"id,omitempty"`
	PaymentType   GetRecurringPaymentType `json:"paymentType,omitempty"`
	AccountNumber int32                   `json:"accountNumber,omitempty"`
	CheckNumber   int32                   `json:"checkNumber,omitempty"`
	RoutingCode   int32                   `json:"routingCode,omitempty"`
	ReceiptNumber int32                   `json:"receiptNumber,omitempty"`
	BankNumber    int32                   `json:"bankNumber,omitempty"`
}
