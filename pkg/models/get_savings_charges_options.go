package models

// GetSavingsChargesOptions struct for GetSavingsChargesOptions
type GetSavingsChargesOptions struct {
	Id                    int32                           `json:"id,omitempty"`
	Name                  string                          `json:"name,omitempty"`
	Active                bool                            `json:"active,omitempty"`
	Penalty               bool                            `json:"penalty,omitempty"`
	Currency              GetChargesCurrencyResponse      `json:"currency,omitempty"`
	Amount                float32                         `json:"amount,omitempty"`
	ChargeTimeType        GetSavingsChargesChargeTimeType `json:"chargeTimeType,omitempty"`
	ChargesAppliesTo      GetChargesAppliesTo             `json:"chargesAppliesTo,omitempty"`
	ChargeCalculationType GetChargesChargeCalculationType `json:"chargeCalculationType,omitempty"`
}
