package models

// GetSelfLoansLoanIdResponse GetSelfLoansLoanIdResponse
type GetSelfLoansLoanIdResponse struct {
	Id                              int64                                       `json:"id,omitempty"`
	AccountNo                       int64                                       `json:"accountNo,omitempty"`
	Status                          GetLoansLoanIdStatus                        `json:"status,omitempty"`
	ClientId                        int32                                       `json:"clientId,omitempty"`
	ClientName                      string                                      `json:"clientName,omitempty"`
	ClientOfficeId                  int32                                       `json:"clientOfficeId,omitempty"`
	LoanProductId                   int32                                       `json:"loanProductId,omitempty"`
	LoanProductName                 string                                      `json:"loanProductName,omitempty"`
	LoanProductDescription          string                                      `json:"loanProductDescription,omitempty"`
	LoanPurposeId                   int32                                       `json:"loanPurposeId,omitempty"`
	LoanPurposeName                 string                                      `json:"loanPurposeName,omitempty"`
	LoanOfficerId                   int32                                       `json:"loanOfficerId,omitempty"`
	LoanOfficerName                 string                                      `json:"loanOfficerName,omitempty"`
	LoanType                        GetLoansLoanIdLoanType                      `json:"loanType,omitempty"`
	Currency                        GetLoansLoanIdCurrency                      `json:"currency,omitempty"`
	Principal                       int64                                       `json:"principal,omitempty"`
	TermFrequency                   int32                                       `json:"termFrequency,omitempty"`
	TermPeriodFrequencyType         GetLoansLoanIdTermPeriodFrequencyType       `json:"termPeriodFrequencyType,omitempty"`
	NumberOfRepayments              int32                                       `json:"numberOfRepayments,omitempty"`
	RepaymentEvery                  int32                                       `json:"repaymentEvery,omitempty"`
	RepaymentFrequencyType          GetLoansLoanIdRepaymentFrequencyType        `json:"repaymentFrequencyType,omitempty"`
	InterestRatePerPeriod           int32                                       `json:"interestRatePerPeriod,omitempty"`
	InterestRateFrequencyType       GetLoansLoanIdInterestRateFrequencyType     `json:"interestRateFrequencyType,omitempty"`
	AnnualInterestRate              int32                                       `json:"annualInterestRate,omitempty"`
	AmortizationType                GetLoansLoanIdAmortizationType              `json:"amortizationType,omitempty"`
	InterestType                    GetLoansLoanIdInterestType                  `json:"interestType,omitempty"`
	InterestCalculationPeriodType   GetLoansLoanIdInterestCalculationPeriodType `json:"interestCalculationPeriodType,omitempty"`
	TransactionProcessingStrategyId int32                                       `json:"transactionProcessingStrategyId,omitempty"`
	Timeline                        GetLoansLoanIdTimeline                      `json:"timeline,omitempty"`
	Summary                         GetLoansLoanIdSummary                       `json:"summary,omitempty"`
}
