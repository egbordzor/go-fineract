package models

// LoanProductConfigurableAttributes struct for LoanProductConfigurableAttributes
type LoanProductConfigurableAttributes struct {
	Id                                        int64       `json:"id,omitempty"`
	LoanProduct                               LoanProduct `json:"loanProduct,omitempty"`
	AmortizationType                          bool        `json:"amortizationType,omitempty"`
	InterestType                              bool        `json:"interestType,omitempty"`
	TransactionProcessingStrategyId           bool        `json:"transactionProcessingStrategyId,omitempty"`
	InterestCalculationPeriodType             bool        `json:"interestCalculationPeriodType,omitempty"`
	InArrearsTolerance                        bool        `json:"inArrearsTolerance,omitempty"`
	RepaymentEvery                            bool        `json:"repaymentEvery,omitempty"`
	GraceOnPrincipalAndInterestPayment        bool        `json:"graceOnPrincipalAndInterestPayment,omitempty"`
	GraceOnArrearsAgeing                      bool        `json:"graceOnArrearsAgeing,omitempty"`
	AmortizationBoolean                       bool        `json:"amortizationBoolean,omitempty"`
	InterestMethodBoolean                     bool        `json:"interestMethodBoolean,omitempty"`
	TransactionProcessingStrategyBoolean      bool        `json:"transactionProcessingStrategyBoolean,omitempty"`
	InterestCalcPeriodBoolean                 bool        `json:"interestCalcPeriodBoolean,omitempty"`
	ArrearsToleranceBoolean                   bool        `json:"arrearsToleranceBoolean,omitempty"`
	RepaymentEveryBoolean                     bool        `json:"repaymentEveryBoolean,omitempty"`
	GraceOnPrincipalAndInterestPaymentBoolean bool        `json:"graceOnPrincipalAndInterestPaymentBoolean,omitempty"`
	GraceOnArrearsAgingBoolean                bool        `json:"graceOnArrearsAgingBoolean,omitempty"`
	New                                       bool        `json:"new,omitempty"`
}
