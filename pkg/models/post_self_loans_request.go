package models

// PostSelfLoansRequest PostSelfLoansRequest
type PostSelfLoansRequest struct {
	DateFormat                      string                          `json:"dateFormat,omitempty"`
	Locale                          string                          `json:"locale,omitempty"`
	ClientId                        int32                           `json:"clientId,omitempty"`
	ProductId                       int32                           `json:"productId,omitempty"`
	Principal                       float64                         `json:"principal,omitempty"`
	LoanTermFrequency               int32                           `json:"loanTermFrequency,omitempty"`
	LoanTermFrequencyType           int32                           `json:"loanTermFrequencyType,omitempty"`
	LoanType                        string                          `json:"loanType,omitempty"`
	NumberOfRepayments              int32                           `json:"numberOfRepayments,omitempty"`
	RepaymentEvery                  int32                           `json:"repaymentEvery,omitempty"`
	RepaymentFrequencyType          int32                           `json:"repaymentFrequencyType,omitempty"`
	InterestRatePerPeriod           int32                           `json:"interestRatePerPeriod,omitempty"`
	AmortizationType                int32                           `json:"amortizationType,omitempty"`
	InterestType                    int32                           `json:"interestType,omitempty"`
	InterestCalculationPeriodType   int32                           `json:"interestCalculationPeriodType,omitempty"`
	TransactionProcessingStrategyId int32                           `json:"transactionProcessingStrategyId,omitempty"`
	ExpectedDisbursementDate        string                          `json:"expectedDisbursementDate,omitempty"`
	SubmittedOnDate                 string                          `json:"submittedOnDate,omitempty"`
	LinkAccountId                   int32                           `json:"linkAccountId,omitempty"`
	FixedEmiAmount                  int32                           `json:"fixedEmiAmount,omitempty"`
	MaxOutstandingLoanBalance       int64                           `json:"maxOutstandingLoanBalance,omitempty"`
	DisbursementData                []PostSelfLoansDisbursementData `json:"disbursementData,omitempty"`
	Datatables                      []PostSelfLoansDatatables       `json:"datatables,omitempty"`
}
