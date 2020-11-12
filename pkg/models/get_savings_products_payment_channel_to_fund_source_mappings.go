package models

// GetSavingsProductsPaymentChannelToFundSourceMappings struct for GetSavingsProductsPaymentChannelToFundSourceMappings
type GetSavingsProductsPaymentChannelToFundSourceMappings struct {
	PaymentType       GetSavingsProductsPaymentType       `json:"paymentType,omitempty"`
	FundSourceAccount GetSavingsProductsFundSourceAccount `json:"fundSourceAccount,omitempty"`
}
