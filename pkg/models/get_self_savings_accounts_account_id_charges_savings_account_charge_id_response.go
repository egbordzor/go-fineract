package models

// GetSelfSavingsAccountsAccountIdChargesSavingsAccountChargeIdResponse GetSelfSavingsAccountsAccountIdChargesSavingsAccountChargeIdResponse
type GetSelfSavingsAccountsAccountIdChargesSavingsAccountChargeIdResponse struct {
	Id                        int32                               `json:"id,omitempty"`
	ChargeId                  int32                               `json:"chargeId,omitempty"`
	Name                      string                              `json:"name,omitempty"`
	ChargeTimeType            GetSelfSavingsChargeTimeType        `json:"chargeTimeType,omitempty"`
	ChargeCalculationType     GetSelfSavingsChargeCalculationType `json:"chargeCalculationType,omitempty"`
	Percentage                float64                             `json:"percentage,omitempty"`
	AmountPercentageAppliedTo float64                             `json:"amountPercentageAppliedTo,omitempty"`
	Currency                  GetSelfSavingsCurrency              `json:"currency,omitempty"`
	Amount                    int32                               `json:"amount,omitempty"`
	AmountPaid                int32                               `json:"amountPaid,omitempty"`
	AmountWaived              int32                               `json:"amountWaived,omitempty"`
	AmountWrittenOff          int32                               `json:"amountWrittenOff,omitempty"`
	AmountOutstanding         int32                               `json:"amountOutstanding,omitempty"`
	AmountOrPercentage        int32                               `json:"amountOrPercentage,omitempty"`
	Penalty                   bool                                `json:"penalty,omitempty"`
}
