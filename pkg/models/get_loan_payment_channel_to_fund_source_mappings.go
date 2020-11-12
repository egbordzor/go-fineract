package models

// GetLoanPaymentChannelToFundSourceMappings struct for GetLoanPaymentChannelToFundSourceMappings
type GetLoanPaymentChannelToFundSourceMappings struct {
	PaymentType       GetLoanPaymentType       `json:"paymentType,omitempty"`
	FundSourceAccount GetLoanFundSourceAccount `json:"fundSourceAccount,omitempty"`
}
