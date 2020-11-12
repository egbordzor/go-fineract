package models

// GetLoansLoanIdOverdueCharges struct for GetLoansLoanIdOverdueCharges
type GetLoansLoanIdOverdueCharges struct {
	Id                    int32                                `json:"id,omitempty"`
	Name                  string                               `json:"name,omitempty"`
	Active                bool                                 `json:"active,omitempty"`
	Penalty               bool                                 `json:"penalty,omitempty"`
	Currency              GetLoanCurrency                      `json:"currency,omitempty"`
	Amount                float32                              `json:"amount,omitempty"`
	ChargeTimeType        GetLoansLoanIdChargeTimeType         `json:"chargeTimeType,omitempty"`
	ChargeAppliesTo       GetLoanChargeTemplateChargeAppliesTo `json:"chargeAppliesTo,omitempty"`
	ChargeCalculationType GetLoansLoanIdChargeCalculationType  `json:"chargeCalculationType,omitempty"`
	ChargePaymentMode     GetLoansLoanIdChargePaymentMode      `json:"chargePaymentMode,omitempty"`
	FeeInterval           int32                                `json:"feeInterval,omitempty"`
	FeeFrequency          GetLoansLoanIdFeeFrequency           `json:"feeFrequency,omitempty"`
}
