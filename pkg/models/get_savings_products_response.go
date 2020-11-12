package models

// GetSavingsProductsResponse GetSavingsProductsResponse
type GetSavingsProductsResponse struct {
	Id                                int32                                               `json:"id,omitempty"`
	Name                              string                                              `json:"name,omitempty"`
	ShortName                         string                                              `json:"shortName,omitempty"`
	Description                       string                                              `json:"description,omitempty"`
	Currency                          GetSavingsCurrency                                  `json:"currency,omitempty"`
	NominalAnnualInterestRate         float32                                             `json:"nominalAnnualInterestRate,omitempty"`
	InterestCompoundingPeriodType     GetSavingsProductsInterestCompoundingPeriodType     `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         GetSavingsProductsInterestPostingPeriodType         `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           GetSavingsProductsInterestCalculationType           `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType GetSavingsProductsInterestCalculationDaysInYearType `json:"interestCalculationDaysInYearType,omitempty"`
	WithdrawalFeeForTransfers         bool                                                `json:"withdrawalFeeForTransfers,omitempty"`
	AccountingRule                    GetSavingsProductsAccountingRule                    `json:"accountingRule,omitempty"`
}
