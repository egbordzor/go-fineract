package models

// GetChargesResponse GetChargesResponse
type GetChargesResponse struct {
	Id                    int64                             `json:"id,omitempty"`
	Name                  string                            `json:"name,omitempty"`
	Active                string                            `json:"active,omitempty"`
	Penalty               string                            `json:"penalty,omitempty"`
	Currency              GetChargesCurrencyResponse        `json:"currency,omitempty"`
	Amount                float32                           `json:"amount,omitempty"`
	ChargeTimeType        GetChargesTimeTypeResponse        `json:"chargeTimeType,omitempty"`
	ChargeAppliesTo       GetChargesAppliesToResponse       `json:"chargeAppliesTo,omitempty"`
	ChargeCalculationType GetChargesCalculationTypeResponse `json:"chargeCalculationType,omitempty"`
	ChargePaymentMode     GetChargesPaymentModeResponse     `json:"chargePaymentMode,omitempty"`
}
