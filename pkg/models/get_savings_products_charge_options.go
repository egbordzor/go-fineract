package models

// GetSavingsProductsChargeOptions struct for GetSavingsProductsChargeOptions
type GetSavingsProductsChargeOptions struct {
	Active                bool                              `json:"active,omitempty"`
	Amount                int64                             `json:"amount,omitempty"`
	ChargeAppliesTo       GetSavingsProductsChargeAppliesTo `json:"chargeAppliesTo,omitempty"`
	ChargeCalculationType GetSavingsChargeCalculationType   `json:"chargeCalculationType,omitempty"`
	ChargePaymentMode     GetSavingsChargePaymentMode       `json:"chargePaymentMode,omitempty"`
	ChargeTimeType        GetSavingsChargeTimeType          `json:"chargeTimeType,omitempty"`
	Currency              GetSavingsCurrency                `json:"currency,omitempty"`
	Id                    int32                             `json:"id,omitempty"`
	Name                  string                            `json:"name,omitempty"`
	Penalty               bool                              `json:"penalty,omitempty"`
}
