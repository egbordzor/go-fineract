package models

// GetLoanProductsResponse GetLoanProductsResponse
type GetLoanProductsResponse struct {
	Id                                          int32                                         `json:"id,omitempty"`
	Name                                        string                                        `json:"name,omitempty"`
	ShortName                                   string                                        `json:"shortName,omitempty"`
	IncludeInBorrowerCycle                      bool                                          `json:"includeInBorrowerCycle,omitempty"`
	UseBorrowerCycle                            bool                                          `json:"useBorrowerCycle,omitempty"`
	StartDate                                   string                                        `json:"startDate,omitempty"`
	EndDate                                     string                                        `json:"endDate,omitempty"`
	Status                                      string                                        `json:"status,omitempty"`
	Currency                                    GetLoanProductsCurrency                       `json:"currency,omitempty"`
	Principal                                   float32                                       `json:"principal,omitempty"`
	MinPrincipal                                float32                                       `json:"minPrincipal,omitempty"`
	MaxPrincipal                                float32                                       `json:"maxPrincipal,omitempty"`
	NumberOfRepayments                          int32                                         `json:"numberOfRepayments,omitempty"`
	MinNumberOfRepayments                       int32                                         `json:"minNumberOfRepayments,omitempty"`
	MaxNumberOfRepayments                       int32                                         `json:"maxNumberOfRepayments,omitempty"`
	RepaymentEvery                              int32                                         `json:"repaymentEvery,omitempty"`
	RepaymentFrequencyType                      GetLoanProductsRepaymentFrequencyType         `json:"repaymentFrequencyType,omitempty"`
	InterestRatePerPeriod                       float32                                       `json:"interestRatePerPeriod,omitempty"`
	InterestRateFrequencyType                   GetLoanProductsInterestRateFrequencyType      `json:"interestRateFrequencyType,omitempty"`
	AnnualInterestRate                          float32                                       `json:"annualInterestRate,omitempty"`
	AmortizationType                            GetLoanProductsAmortizationType               `json:"amortizationType,omitempty"`
	InterestType                                GetLoanProductsInterestType                   `json:"interestType,omitempty"`
	InterestCalculationPeriodType               GetLoansProductsInterestCalculationPeriodType `json:"interestCalculationPeriodType,omitempty"`
	TransactionProcessingStrategyId             int32                                         `json:"transactionProcessingStrategyId,omitempty"`
	TransactionProcessingStrategyName           string                                        `json:"transactionProcessingStrategyName,omitempty"`
	PrincipalVariationsForBorrowerCycle         []int32                                       `json:"principalVariationsForBorrowerCycle,omitempty"`
	InterestRateVariationsForBorrowerCycle      []int32                                       `json:"interestRateVariationsForBorrowerCycle,omitempty"`
	NumberOfRepaymentVariationsForBorrowerCycle []int32                                       `json:"numberOfRepaymentVariationsForBorrowerCycle,omitempty"`
	DaysInMonthType                             GetLoansProductsDaysInMonthType               `json:"daysInMonthType,omitempty"`
	DaysInYearType                              GetLoansProductsDaysInYearType                `json:"daysInYearType,omitempty"`
	IsInterestRecalculationEnabled              bool                                          `json:"isInterestRecalculationEnabled,omitempty"`
	InterestRecalculationData                   GetLoanProductsInterestRecalculationData      `json:"interestRecalculationData,omitempty"`
	AccountingRule                              GetLoanProductsAccountingRule                 `json:"accountingRule,omitempty"`
	PrincipalThresholdForLastInstalment         int32                                         `json:"principalThresholdForLastInstalment,omitempty"`
}
