package models

// PostLoansRequest PostLoansRequest
type PostLoansRequest struct {
	DateFormat                      string  `json:"dateFormat,omitempty"`
	Locale                          string  `json:"locale,omitempty"`
	ProductId                       int32   `json:"productId,omitempty"`
	Principal                       float64 `json:"principal,omitempty"`
	LoanTermFrequency               int32   `json:"loanTermFrequency,omitempty"`
	LoanTermFrequencyType           int32   `json:"loanTermFrequencyType,omitempty"`
	NumberOfRepayments              int32   `json:"numberOfRepayments,omitempty"`
	RepaymentEvery                  int32   `json:"repaymentEvery,omitempty"`
	RepaymentFrequencyType          int32   `json:"repaymentFrequencyType,omitempty"`
	InterestRatePerPeriod           int32   `json:"interestRatePerPeriod,omitempty"`
	AmortizationType                int32   `json:"amortizationType,omitempty"`
	InterestType                    int32   `json:"interestType,omitempty"`
	InterestCalculationPeriodType   int32   `json:"interestCalculationPeriodType,omitempty"`
	ExpectedDisbursementDate        string  `json:"expectedDisbursementDate,omitempty"`
	TransactionProcessingStrategyId int32   `json:"transactionProcessingStrategyId,omitempty"`
}
