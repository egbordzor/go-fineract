package models

// GetFixedDepositProductsProductIdResponse GetFixedDepositProductsProductIdResponse
type GetFixedDepositProductsProductIdResponse struct {
	Id                                int32                                                            `json:"id,omitempty"`
	Name                              string                                                           `json:"name,omitempty"`
	ShortName                         string                                                           `json:"shortName,omitempty"`
	Description                       string                                                           `json:"description,omitempty"`
	Currency                          GetFixedDepositProductsProductIdCurrency                         `json:"currency,omitempty"`
	InterestCompoundingPeriodType     GetFixedDepositProductsProductIdInterestCompoundingPeriodType    `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         GetFixedDepositProductsInterestPostingPeriodType                 `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           GetFixedDepositProductsInterestCalculationType                   `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType GetFixedDepositProductsInterestCalculationDaysInYearType         `json:"interestCalculationDaysInYearType,omitempty"`
	AccountingMappings                GetFixedDepositProductsProductIdAccountingMappings               `json:"accountingMappings,omitempty"`
	FeeToIncomeAccountMappings        []GetFixedDepositProductsProductIdFeeToIncomeAccountMappings     `json:"feeToIncomeAccountMappings,omitempty"`
	PenaltyToIncomeAccountMappings    []GetFixedDepositProductsProductIdPenaltyToIncomeAccountMappings `json:"penaltyToIncomeAccountMappings,omitempty"`
	PreClosurePenalApplicable         bool                                                             `json:"preClosurePenalApplicable,omitempty"`
	PreClosurePenalInterest           float64                                                          `json:"preClosurePenalInterest,omitempty"`
	PreClosurePenalInterestOnType     GetFixedDepositProductsProductIdPreClosurePenalInterestOnType    `json:"preClosurePenalInterestOnType,omitempty"`
	MinDepositTerm                    int32                                                            `json:"minDepositTerm,omitempty"`
	MinDepositTermType                GetFixedDepositProductsProductIdMinDepositTermType               `json:"minDepositTermType,omitempty"`
	MaxDepositTerm                    int32                                                            `json:"maxDepositTerm,omitempty"`
	MaxDepositTermType                GetFixedDepositProductsProductIdMaxDepositTermType               `json:"maxDepositTermType,omitempty"`
	ActiveChart                       GetFixedDepositProductsProductIdActiveChart                      `json:"activeChart,omitempty"`
}
