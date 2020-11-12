package models

// GetAccountsCharges struct for GetAccountsCharges
type GetAccountsCharges struct {
	Id                        int32                            `json:"id,omitempty"`
	ChargeId                  int32                            `json:"chargeId,omitempty"`
	AccountId                 int32                            `json:"accountId,omitempty"`
	Name                      string                           `json:"name,omitempty"`
	ChargeTimeType            GetAccountsChargeTimeType        `json:"chargeTimeType,omitempty"`
	ChargeCalculationType     GetAccountsChargeCalculationType `json:"chargeCalculationType,omitempty"`
	Percentage                float64                          `json:"percentage,omitempty"`
	AmountPercentageAppliedTo float64                          `json:"amountPercentageAppliedTo,omitempty"`
	Currency                  GetAccountsChargesCurrency       `json:"currency,omitempty"`
	Amount                    float32                          `json:"amount,omitempty"`
	AmountPaid                float32                          `json:"amountPaid,omitempty"`
	AmountWaived              float32                          `json:"amountWaived,omitempty"`
	AmountWrittenOff          float32                          `json:"amountWrittenOff,omitempty"`
	AmountOutstanding         float32                          `json:"amountOutstanding,omitempty"`
	AmountOrPercentage        float32                          `json:"amountOrPercentage,omitempty"`
	IsActive                  bool                             `json:"isActive,omitempty"`
}
