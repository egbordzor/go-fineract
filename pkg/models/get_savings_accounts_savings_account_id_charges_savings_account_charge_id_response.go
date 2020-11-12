package models

// GetSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse GetSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse
type GetSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse struct {
	Id                        int32                           `json:"id,omitempty"`
	ChargeId                  int32                           `json:"chargeId,omitempty"`
	Name                      string                          `json:"name,omitempty"`
	ChargeTimeType            GetChargesChargeTimeType        `json:"chargeTimeType,omitempty"`
	ChargeCalculationType     GetChargesChargeCalculationType `json:"chargeCalculationType,omitempty"`
	Percentage                float64                         `json:"percentage,omitempty"`
	AmountPercentageAppliedTo float64                         `json:"amountPercentageAppliedTo,omitempty"`
	Currency                  GetChargesCurrencyResponse      `json:"currency,omitempty"`
	Amount                    float32                         `json:"amount,omitempty"`
	AmountPaid                float32                         `json:"amountPaid,omitempty"`
	AmountWaived              float32                         `json:"amountWaived,omitempty"`
	AmountWrittenOff          float32                         `json:"amountWrittenOff,omitempty"`
	AmountOutstanding         float32                         `json:"amountOutstanding,omitempty"`
	AmountOrPercentage        float32                         `json:"amountOrPercentage,omitempty"`
	Penalty                   bool                            `json:"penalty,omitempty"`
}
