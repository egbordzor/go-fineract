package models

// GetLoanProductsTemplateResponse GetLoanProductsTemplateResponse
type GetLoanProductsTemplateResponse struct {
	IncludeInBorrowerCycle                       bool                                                           `json:"includeInBorrowerCycle,omitempty"`
	UseBorrowerCycle                             bool                                                           `json:"useBorrowerCycle,omitempty"`
	Currency                                     GetLoanProductsTemplateCurrency                                `json:"currency,omitempty"`
	RepaymentFrequencyType                       GetLoanProductsRepaymentTemplateFrequencyType                  `json:"repaymentFrequencyType,omitempty"`
	InterestRateFrequencyType                    GetLoanProductsInterestRateTemplateFrequencyType               `json:"interestRateFrequencyType,omitempty"`
	AmortizationType                             GetLoanProductsAmortizationType                                `json:"amortizationType,omitempty"`
	InterestType                                 GetLoanProductsInterestTemplateType                            `json:"interestType,omitempty"`
	InterestCalculationPeriodType                GetLoansProductsInterestCalculationPeriodType                  `json:"interestCalculationPeriodType,omitempty"`
	PrincipalVariationsForBorrowerCycle          []int32                                                        `json:"principalVariationsForBorrowerCycle,omitempty"`
	InterestRateVariationsForBorrowerCycle       []int32                                                        `json:"interestRateVariationsForBorrowerCycle,omitempty"`
	NumberOfRepaymentVariationsForBorrowerCycle  []int32                                                        `json:"numberOfRepaymentVariationsForBorrowerCycle,omitempty"`
	AccountingRule                               GetLoanProductsAccountingRule                                  `json:"accountingRule,omitempty"`
	DaysInMonthType                              GetLoansProductsDaysInMonthTemplateType                        `json:"daysInMonthType,omitempty"`
	DaysInYearType                               GetLoanProductsDaysInYearTemplateType                          `json:"daysInYearType,omitempty"`
	IsInterestRecalculationEnabled               bool                                                           `json:"isInterestRecalculationEnabled,omitempty"`
	InterestRecalculationData                    GetLoanProductsInterestRecalculationTemplateData               `json:"interestRecalculationData,omitempty"`
	PaymentTypeOptions                           []GetLoanProductsPaymentTypeOptions                            `json:"paymentTypeOptions,omitempty"`
	CurrencyOptions                              []GetLoanProductsCurrencyOptions                               `json:"currencyOptions,omitempty"`
	RepaymentFrequencyTypeOptions                []GetLoanProductsRepaymentTemplateFrequencyType                `json:"repaymentFrequencyTypeOptions,omitempty"`
	PreClosureInterestCalculationStrategyOptions []GetLoanProductsPreClosureInterestCalculationStrategy         `json:"preClosureInterestCalculationStrategyOptions,omitempty"`
	InterestRateFrequencyTypeOptions             []GetLoanProductsInterestRateTemplateFrequencyType             `json:"interestRateFrequencyTypeOptions,omitempty"`
	AmortizationTypeOptions                      []GetLoanProductsAmortizationType                              `json:"amortizationTypeOptions,omitempty"`
	InterestTypeOptions                          []GetLoanProductsInterestTemplateType                          `json:"interestTypeOptions,omitempty"`
	InterestCalculationPeriodTypeOptions         []GetLoansProductsInterestCalculationPeriodType                `json:"interestCalculationPeriodTypeOptions,omitempty"`
	TransactionProcessingStrategyOptions         []GetLoanProductsTransactionProcessingStrategyOptions          `json:"transactionProcessingStrategyOptions,omitempty"`
	ChargeOptions                                []GetLoanProductsChargeOptions                                 `json:"chargeOptions,omitempty"`
	AccountingRuleOptions                        []GetLoanProductsAccountingRule                                `json:"accountingRuleOptions,omitempty"`
	AccountingMappingOptions                     GetLoanProductsAccountingMappingOptions                        `json:"accountingMappingOptions,omitempty"`
	ValueConditionTypeOptions                    []GetLoanProductsValueConditionTypeOptions                     `json:"valueConditionTypeOptions,omitempty"`
	DaysInMonthTypeOptions                       []GetLoansProductsDaysInMonthTemplateType                      `json:"daysInMonthTypeOptions,omitempty"`
	DaysInYearTypeOptions                        []GetLoanProductsInterestTemplateType                          `json:"daysInYearTypeOptions,omitempty"`
	InterestRecalculationCompoundingTypeOptions  []GetLoanProductsInterestRecalculationCompoundingType          `json:"interestRecalculationCompoundingTypeOptions,omitempty"`
	RescheduleStrategyTypeOptions                []GetLoanProductsRescheduleStrategyType                        `json:"rescheduleStrategyTypeOptions,omitempty"`
	InterestRecalculationFrequencyTypeOptions    []GetLoanProductsInterestRecalculationCompoundingFrequencyType `json:"interestRecalculationFrequencyTypeOptions,omitempty"`
}
