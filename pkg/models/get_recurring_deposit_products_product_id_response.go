package models

// GetRecurringDepositProductsProductIdResponse GetRecurringDepositProductsProductIdResponse
type GetRecurringDepositProductsProductIdResponse struct {
	Id                                int32                                                                `json:"id,omitempty"`
	Name                              string                                                               `json:"name,omitempty"`
	ShortName                         string                                                               `json:"shortName,omitempty"`
	Description                       string                                                               `json:"description,omitempty"`
	Currency                          GetRecurringDepositProductsProductIdCurrency                         `json:"currency,omitempty"`
	InterestCompoundingPeriodType     GetRecurringDepositProductsProductIdInterestCompoundingPeriodType    `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         GetRecurringDepositProductsInterestPostingPeriodType                 `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           GetRecurringDepositProductsInterestCalculationType                   `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType GetRecurringDepositProductsInterestCalculationDaysInYearType         `json:"interestCalculationDaysInYearType,omitempty"`
	AccountingMappings                GetRecurringDepositProductsProductIdAccountingMappings               `json:"accountingMappings,omitempty"`
	FeeToIncomeAccountMappings        []GetRecurringDepositProductsProductIdFeeToIncomeAccountMappings     `json:"feeToIncomeAccountMappings,omitempty"`
	PenaltyToIncomeAccountMappings    []GetRecurringDepositProductsProductIdPenaltyToIncomeAccountMappings `json:"penaltyToIncomeAccountMappings,omitempty"`
	RecurringDepositFrequency         int32                                                                `json:"recurringDepositFrequency,omitempty"`
	RecurringDepositFrequencyType     GetRecurringDepositProductsRecurringDepositFrequencyType             `json:"recurringDepositFrequencyType,omitempty"`
	PreClosurePenalApplicable         bool                                                                 `json:"preClosurePenalApplicable,omitempty"`
	PreClosurePenalInterest           float64                                                              `json:"preClosurePenalInterest,omitempty"`
	PreClosurePenalInterestOnType     GetRecurringDepositProductsProductIdPreClosurePenalInterestOnType    `json:"preClosurePenalInterestOnType,omitempty"`
	MinDepositTerm                    int32                                                                `json:"minDepositTerm,omitempty"`
	MinDepositTermType                GetRecurringDepositProductsProductIdMinDepositTermType               `json:"minDepositTermType,omitempty"`
	MaxDepositTerm                    int32                                                                `json:"maxDepositTerm,omitempty"`
	MaxDepositTermType                GetRecurringDepositProductsProductIdMaxDepositTermType               `json:"maxDepositTermType,omitempty"`
	ActiveChart                       GetRecurringDepositProductsProductIdActiveChart                      `json:"activeChart,omitempty"`
}
