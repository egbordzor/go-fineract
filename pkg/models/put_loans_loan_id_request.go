package models

// PutLoansLoanIdRequest PutLoansLoanIdRequest
type PutLoansLoanIdRequest struct {
	Locale                          string `json:"locale,omitempty"`
	DateFormat                      string `json:"dateFormat,omitempty"`
	ProductId                       int32  `json:"productId,omitempty"`
	Principal                       int64  `json:"principal,omitempty"`
	LoanTermFrequency               int32  `json:"loanTermFrequency,omitempty"`
	LoanTermFrequencyType           int32  `json:"loanTermFrequencyType,omitempty"`
	NumberOfRepayments              int32  `json:"numberOfRepayments,omitempty"`
	RepaymentEvery                  int32  `json:"repaymentEvery,omitempty"`
	RepaymentFrequencyType          int32  `json:"repaymentFrequencyType,omitempty"`
	InterestRatePerPeriod           int32  `json:"interestRatePerPeriod,omitempty"`
	InterestType                    int32  `json:"interestType,omitempty"`
	InterestCalculationPeriodType   int32  `json:"interestCalculationPeriodType,omitempty"`
	AmortizationType                int32  `json:"amortizationType,omitempty"`
	ExpectedDisbursementDate        string `json:"expectedDisbursementDate,omitempty"`
	TransactionProcessingStrategyId int32  `json:"transactionProcessingStrategyId,omitempty"`
}
