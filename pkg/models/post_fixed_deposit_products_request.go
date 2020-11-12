package models

// PostFixedDepositProductsRequest PostFixedDepositProductsRequest
type PostFixedDepositProductsRequest struct {
	Name                              string                           `json:"name,omitempty"`
	ShortName                         string                           `json:"shortName,omitempty"`
	Description                       string                           `json:"description,omitempty"`
	CurrencyCode                      string                           `json:"currencyCode,omitempty"`
	DigitsAfterDecimal                int32                            `json:"digitsAfterDecimal,omitempty"`
	InMultiplesOf                     int32                            `json:"inMultiplesOf,omitempty"`
	Locale                            string                           `json:"locale,omitempty"`
	InterestCompoundingPeriodType     int32                            `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         int32                            `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           int32                            `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType int32                            `json:"interestCalculationDaysInYearType,omitempty"`
	AccountingRule                    int32                            `json:"accountingRule,omitempty"`
	PreClosurePenalApplicable         bool                             `json:"preClosurePenalApplicable,omitempty"`
	PreClosurePenalInterest           float64                          `json:"preClosurePenalInterest,omitempty"`
	PreClosurePenalInterestOnTypeId   int32                            `json:"preClosurePenalInterestOnTypeId,omitempty"`
	MinDepositTerm                    int32                            `json:"minDepositTerm,omitempty"`
	MinDepositTermTypeId              int32                            `json:"minDepositTermTypeId,omitempty"`
	MaxDepositTerm                    int32                            `json:"maxDepositTerm,omitempty"`
	MaxDepositTermTypeId              int32                            `json:"maxDepositTermTypeId,omitempty"`
	Charts                            []PostFixedDepositProductsCharts `json:"charts,omitempty"`
}
