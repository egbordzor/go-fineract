package models

// GetLoanProductsProductIdResponse GetLoanProductsProductIdResponse
type GetLoanProductsProductIdResponse struct {
	Id                                          int32                                                `json:"id,omitempty"`
	Name                                        string                                               `json:"name,omitempty"`
	ShortName                                   string                                               `json:"shortName,omitempty"`
	IncludeInBorrowerCycle                      bool                                                 `json:"includeInBorrowerCycle,omitempty"`
	UseBorrowerCycle                            bool                                                 `json:"useBorrowerCycle,omitempty"`
	Status                                      string                                               `json:"status,omitempty"`
	Currency                                    GetLoanProductsCurrency                              `json:"currency,omitempty"`
	Principal                                   float32                                              `json:"principal,omitempty"`
	MinPrincipal                                float32                                              `json:"minPrincipal,omitempty"`
	MaxPrincipal                                float32                                              `json:"maxPrincipal,omitempty"`
	NumberOfRepayments                          int32                                                `json:"numberOfRepayments,omitempty"`
	RepaymentEvery                              int32                                                `json:"repaymentEvery,omitempty"`
	RepaymentFrequencyType                      GetLoanProductsRepaymentFrequencyType                `json:"repaymentFrequencyType,omitempty"`
	InterestRatePerPeriod                       float32                                              `json:"interestRatePerPeriod,omitempty"`
	InterestRateFrequencyType                   GetLoanProductsInterestRateFrequencyType             `json:"interestRateFrequencyType,omitempty"`
	AnnualInterestRate                          float32                                              `json:"annualInterestRate,omitempty"`
	AmortizationType                            GetLoanProductsAmortizationType                      `json:"amortizationType,omitempty"`
	InterestType                                GetLoanProductsInterestTemplateType                  `json:"interestType,omitempty"`
	InterestCalculationPeriodType               GetLoansProductsInterestCalculationPeriodType        `json:"interestCalculationPeriodType,omitempty"`
	TransactionProcessingStrategyId             int32                                                `json:"transactionProcessingStrategyId,omitempty"`
	TransactionProcessingStrategyName           string                                               `json:"transactionProcessingStrategyName,omitempty"`
	Charges                                     []int32                                              `json:"charges,omitempty"`
	ProductsPrincipalVariationsForBorrowerCycle []GetLoanProductsPrincipalVariationsForBorrowerCycle `json:"productsPrincipalVariationsForBorrowerCycle,omitempty"`
	InterestRateVariationsForBorrowerCycle      []int32                                              `json:"interestRateVariationsForBorrowerCycle,omitempty"`
	NumberOfRepaymentVariationsForBorrowerCycle []int32                                              `json:"numberOfRepaymentVariationsForBorrowerCycle,omitempty"`
	AccountingRule                              GetLoanProductsAccountingRule                        `json:"accountingRule,omitempty"`
	AccountingMappings                          GetLoanAccountingMappings                            `json:"accountingMappings,omitempty"`
	PaymentChannelToFundSourceMappings          []GetLoanPaymentChannelToFundSourceMappings          `json:"paymentChannelToFundSourceMappings,omitempty"`
	FeeToIncomeAccountMappings                  []GetLoanFeeToIncomeAccountMappings                  `json:"feeToIncomeAccountMappings,omitempty"`
	MultiDisburseLoan                           bool                                                 `json:"multiDisburseLoan,omitempty"`
	MaxTrancheCount                             int32                                                `json:"maxTrancheCount,omitempty"`
	OutstandingLoanBalance                      float32                                              `json:"outstandingLoanBalance,omitempty"`
	OverdueDaysForNPA                           int32                                                `json:"overdueDaysForNPA,omitempty"`
	PrincipalThresholdForLastInstalment         int32                                                `json:"principalThresholdForLastInstalment,omitempty"`
}
