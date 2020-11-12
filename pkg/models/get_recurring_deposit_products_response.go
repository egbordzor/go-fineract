package models

// GetRecurringDepositProductsResponse GetRecurringDepositProductsResponse
type GetRecurringDepositProductsResponse struct {
	Id                                int32                                                        `json:"id,omitempty"`
	Name                              string                                                       `json:"name,omitempty"`
	ShortName                         string                                                       `json:"shortName,omitempty"`
	Description                       string                                                       `json:"description,omitempty"`
	Currency                          GetRecurringDepositProductsCurrency                          `json:"currency,omitempty"`
	RecurringDepositFrequency         int32                                                        `json:"recurringDepositFrequency,omitempty"`
	RecurringDepositFrequencyType     GetRecurringDepositProductsRecurringDepositFrequencyType     `json:"recurringDepositFrequencyType,omitempty"`
	PreClosurePenalApplicable         bool                                                         `json:"preClosurePenalApplicable,omitempty"`
	MinDepositTerm                    int32                                                        `json:"minDepositTerm,omitempty"`
	MaxDepositTerm                    int32                                                        `json:"maxDepositTerm,omitempty"`
	MinDepositTermType                GetRecurringDepositProductsMinDepositTermType                `json:"minDepositTermType,omitempty"`
	MaxDepositTermType                GetRecurringDepositProductsMaxDepositTermType                `json:"maxDepositTermType,omitempty"`
	NominalAnnualInterestRate         float64                                                      `json:"nominalAnnualInterestRate,omitempty"`
	InterestCompoundingPeriodType     GetRecurringDepositProductsInterestCompoundingPeriodType     `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         GetRecurringDepositProductsInterestPostingPeriodType         `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           GetRecurringDepositProductsInterestCalculationType           `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType GetRecurringDepositProductsInterestCalculationDaysInYearType `json:"interestCalculationDaysInYearType,omitempty"`
	AccountingRule                    GetRecurringDepositProductsAccountingRule                    `json:"accountingRule,omitempty"`
}
