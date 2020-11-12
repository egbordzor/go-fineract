package models

// GetShareAccountsChargeOptions struct for GetShareAccountsChargeOptions
type GetShareAccountsChargeOptions struct {
	Id              int32                                 `json:"id,omitempty"`
	Name            string                                `json:"name,omitempty"`
	Active          bool                                  `json:"active,omitempty"`
	Penalty         bool                                  `json:"penalty,omitempty"`
	Currency        GetShareAccountsCurrency              `json:"currency,omitempty"`
	Amount          int32                                 `json:"amount,omitempty"`
	ChargeTimeType  GetShareAccountsChargeTimeType        `json:"chargeTimeType,omitempty"`
	ChargeAppliesTo GetShareAccountsChargeAppliesTo       `json:"chargeAppliesTo,omitempty"`
	CalculationType GetShareAccountsChargeCalculationType `json:"calculationType,omitempty"`
	PaymentMode     GetShareAccountsChargePaymentMode     `json:"paymentMode,omitempty"`
}
