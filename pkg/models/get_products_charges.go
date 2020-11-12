package models

// GetProductsCharges struct for GetProductsCharges
type GetProductsCharges struct {
	Id                    int32                    `json:"id,omitempty"`
	Name                  string                   `json:"name,omitempty"`
	Active                bool                     `json:"active,omitempty"`
	Penalty               bool                     `json:"penalty,omitempty"`
	Currency              GetChargesCurrency       `json:"currency,omitempty"`
	Amount                int32                    `json:"amount,omitempty"`
	ChargeTimeType        GetChargeTimeType        `json:"chargeTimeType,omitempty"`
	ChargeAppliesTo       GetChargeAppliesTo       `json:"chargeAppliesTo,omitempty"`
	ChargeCalculationType GetChargeCalculationType `json:"chargeCalculationType,omitempty"`
	ChargePaymentMode     GetChargePaymentMode     `json:"chargePaymentMode,omitempty"`
}
