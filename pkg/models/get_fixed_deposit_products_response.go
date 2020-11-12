package models

// GetFixedDepositProductsResponse GetFixedDepositProductsResponse
type GetFixedDepositProductsResponse struct {
	Id                                int32                                                    `json:"id,omitempty"`
	Name                              string                                                   `json:"name,omitempty"`
	ShortName                         string                                                   `json:"shortName,omitempty"`
	Description                       string                                                   `json:"description,omitempty"`
	Currency                          GetFixedDepositProductsCurrency                          `json:"currency,omitempty"`
	PreClosurePenalApplicable         bool                                                     `json:"preClosurePenalApplicable,omitempty"`
	MinDepositTerm                    int32                                                    `json:"minDepositTerm,omitempty"`
	MaxDepositTerm                    int32                                                    `json:"maxDepositTerm,omitempty"`
	MinDepositTermType                GetFixedDepositProductsMinDepositTermType                `json:"minDepositTermType,omitempty"`
	MaxDepositTermType                GetFixedDepositProductsMaxDepositTermType                `json:"maxDepositTermType,omitempty"`
	NominalAnnualInterestRate         float64                                                  `json:"nominalAnnualInterestRate,omitempty"`
	InterestCompoundingPeriodType     GetFixedDepositProductsInterestCompoundingPeriodType     `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         GetFixedDepositProductsInterestPostingPeriodType         `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           GetFixedDepositProductsInterestCalculationType           `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType GetFixedDepositProductsInterestCalculationDaysInYearType `json:"interestCalculationDaysInYearType,omitempty"`
	AccountingRule                    GetFixedDepositProductsAccountingRule                    `json:"accountingRule,omitempty"`
}
