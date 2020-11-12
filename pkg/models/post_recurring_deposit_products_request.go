package models

// PostRecurringDepositProductsRequest PostRecurringDepositProductsRequest
type PostRecurringDepositProductsRequest struct {
	Name                              string                               `json:"name,omitempty"`
	ShortName                         string                               `json:"shortName,omitempty"`
	Description                       string                               `json:"description,omitempty"`
	CurrencyCode                      string                               `json:"currencyCode,omitempty"`
	DigitsAfterDecimal                int32                                `json:"digitsAfterDecimal,omitempty"`
	InMultiplesOf                     int32                                `json:"inMultiplesOf,omitempty"`
	Locale                            string                               `json:"locale,omitempty"`
	InterestCompoundingPeriodType     int32                                `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         int32                                `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           int32                                `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType int32                                `json:"interestCalculationDaysInYearType,omitempty"`
	AccountingRule                    int32                                `json:"accountingRule,omitempty"`
	RecurringDepositFrequency         int32                                `json:"recurringDepositFrequency,omitempty"`
	RecurringDepositFrequencyTypeId   int32                                `json:"recurringDepositFrequencyTypeId,omitempty"`
	PreClosurePenalApplicable         bool                                 `json:"preClosurePenalApplicable,omitempty"`
	PreClosurePenalInterest           float64                              `json:"preClosurePenalInterest,omitempty"`
	PreClosurePenalInterestOnTypeId   int32                                `json:"preClosurePenalInterestOnTypeId,omitempty"`
	MinDepositTerm                    int32                                `json:"minDepositTerm,omitempty"`
	MinDepositTermTypeId              int32                                `json:"minDepositTermTypeId,omitempty"`
	MaxDepositTerm                    int32                                `json:"maxDepositTerm,omitempty"`
	MaxDepositTermTypeId              int32                                `json:"maxDepositTermTypeId,omitempty"`
	DepositAmount                     int64                                `json:"depositAmount,omitempty"`
	MinDepositAmount                  int64                                `json:"minDepositAmount,omitempty"`
	MaxDepositAmount                  int64                                `json:"maxDepositAmount,omitempty"`
	Charts                            []PostRecurringDepositProductsCharts `json:"charts,omitempty"`
}
