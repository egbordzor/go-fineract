package models

// GetLoanProductsInterestRecalculationTemplateData struct for GetLoanProductsInterestRecalculationTemplateData
type GetLoanProductsInterestRecalculationTemplateData struct {
	InterestRecalculationCompoundingType  GetLoanProductsInterestRecalculationCompoundingType  `json:"interestRecalculationCompoundingType,omitempty"`
	RescheduleStrategyType                GetLoanProductsRescheduleStrategyType                `json:"rescheduleStrategyType,omitempty"`
	PreClosureInterestCalculationStrategy GetLoanProductsPreClosureInterestCalculationStrategy `json:"preClosureInterestCalculationStrategy,omitempty"`
}
