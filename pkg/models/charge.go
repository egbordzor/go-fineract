package models

// Charge struct for Charge
type Charge struct {
	Id                                  int64               `json:"id,omitempty"`
	Name                                string              `json:"name,omitempty"`
	Amount                              float32             `json:"amount,omitempty"`
	CurrencyCode                        string              `json:"currencyCode,omitempty"`
	ChargeTimeType                      int32               `json:"chargeTimeType,omitempty"`
	ChargeCalculation                   int32               `json:"chargeCalculation,omitempty"`
	ChargePaymentMode                   int32               `json:"chargePaymentMode,omitempty"`
	FeeInterval                         int32               `json:"feeInterval,omitempty"`
	Penalty                             bool                `json:"penalty,omitempty"`
	Active                              bool                `json:"active,omitempty"`
	Deleted                             bool                `json:"deleted,omitempty"`
	MinCap                              float32             `json:"minCap,omitempty"`
	MaxCap                              float32             `json:"maxCap,omitempty"`
	Account                             GlAccount           `json:"account,omitempty"`
	TaxGroup                            TaxGroup            `json:"taxGroup,omitempty"`
	SavingsCharge                       bool                `json:"savingsCharge,omitempty"`
	FeeOnMonthDay                       ChargeFeeOnMonthDay `json:"feeOnMonthDay,omitempty"`
	AnnualFee                           bool                `json:"annualFee,omitempty"`
	MonthlyFee                          bool                `json:"monthlyFee,omitempty"`
	LoanCharge                          bool                `json:"loanCharge,omitempty"`
	ClientCharge                        bool                `json:"clientCharge,omitempty"`
	OverdueInstallment                  bool                `json:"overdueInstallment,omitempty"`
	AllowedLoanChargeTime               bool                `json:"allowedLoanChargeTime,omitempty"`
	AllowedClientChargeTime             bool                `json:"allowedClientChargeTime,omitempty"`
	AllowedSavingsChargeTime            bool                `json:"allowedSavingsChargeTime,omitempty"`
	DisbursementCharge                  bool                `json:"disbursementCharge,omitempty"`
	AllowedSavingsChargeCalculationType bool                `json:"allowedSavingsChargeCalculationType,omitempty"`
	AllowedClientChargeCalculationType  bool                `json:"allowedClientChargeCalculationType,omitempty"`
	PercentageOfApprovedAmount          bool                `json:"percentageOfApprovedAmount,omitempty"`
	PercentageOfDisbursementAmount      bool                `json:"percentageOfDisbursementAmount,omitempty"`
	New                                 bool                `json:"new,omitempty"`
}
