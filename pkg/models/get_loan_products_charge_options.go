package models

// GetLoanProductsChargeOptions struct for GetLoanProductsChargeOptions
type GetLoanProductsChargeOptions struct {
	Id                    int32                          `json:"id,omitempty"`
	Name                  string                         `json:"name,omitempty"`
	Active                bool                           `json:"active,omitempty"`
	Penalty               bool                           `json:"penalty,omitempty"`
	Currency              GetLoanProductsCurrencyOptions `json:"currency,omitempty"`
	Amount                int64                          `json:"amount,omitempty"`
	ChargeTimeType        GetLoanChargeTimeType          `json:"chargeTimeType,omitempty"`
	ChargeAppliesTo       GetLoanProductsChargeAppliesTo `json:"chargeAppliesTo,omitempty"`
	ChargeCalculationType GetLoanChargeCalculationType   `json:"chargeCalculationType,omitempty"`
	ChargePaymentMode     GetLoansChargePaymentMode      `json:"chargePaymentMode,omitempty"`
}
