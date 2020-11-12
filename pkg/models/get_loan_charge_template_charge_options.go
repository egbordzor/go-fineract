package models

// GetLoanChargeTemplateChargeOptions struct for GetLoanChargeTemplateChargeOptions
type GetLoanChargeTemplateChargeOptions struct {
	Id                    int32                                `json:"id,omitempty"`
	Name                  string                               `json:"name,omitempty"`
	Active                bool                                 `json:"active,omitempty"`
	Penalty               bool                                 `json:"penalty,omitempty"`
	Currency              GetLoanChargeCurrency                `json:"currency,omitempty"`
	Amount                float32                              `json:"amount,omitempty"`
	ChargeTimeType        GetLoanChargeTemplateChargeTimeType  `json:"chargeTimeType,omitempty"`
	ChargeAppliesTo       GetLoanChargeTemplateChargeAppliesTo `json:"chargeAppliesTo,omitempty"`
	ChargeCalculationType GetLoanChargeCalculationType         `json:"chargeCalculationType,omitempty"`
}
