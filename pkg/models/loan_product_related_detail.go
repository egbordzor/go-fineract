package models

// LoanProductRelatedDetail struct for LoanProductRelatedDetail
type LoanProductRelatedDetail struct {
	Currency                              MonetaryCurrency `json:"currency,omitempty"`
	Principal                             Money            `json:"principal,omitempty"`
	NominalInterestRatePerPeriod          float32          `json:"nominalInterestRatePerPeriod,omitempty"`
	InterestPeriodFrequencyType           string           `json:"interestPeriodFrequencyType,omitempty"`
	AnnualNominalInterestRate             float32          `json:"annualNominalInterestRate,omitempty"`
	InterestMethod                        string           `json:"interestMethod,omitempty"`
	InterestCalculationPeriodMethod       string           `json:"interestCalculationPeriodMethod,omitempty"`
	AllowPartialPeriodInterestCalcualtion bool             `json:"allowPartialPeriodInterestCalcualtion,omitempty"`
	RepayEvery                            int32            `json:"repayEvery,omitempty"`
	RepaymentPeriodFrequencyType          string           `json:"repaymentPeriodFrequencyType,omitempty"`
	NumberOfRepayments                    int32            `json:"numberOfRepayments,omitempty"`
	GraceOnPrincipalPayment               int32            `json:"graceOnPrincipalPayment,omitempty"`
	GraceOnInterestPayment                int32            `json:"graceOnInterestPayment,omitempty"`
	AmortizationMethod                    string           `json:"amortizationMethod,omitempty"`
	InArrearsTolerance                    Money            `json:"inArrearsTolerance,omitempty"`
	GraceOnArrearsAgeing                  int32            `json:"graceOnArrearsAgeing,omitempty"`
	EqualAmortization                     bool             `json:"equalAmortization,omitempty"`
	InterestRecalculationEnabled          bool             `json:"interestRecalculationEnabled,omitempty"`
	GraceOnDueDate                        int32            `json:"graceOnDueDate,omitempty"`
	ArrearsTolerance                      float32          `json:"arrearsTolerance,omitempty"`
}
