package models

// GetLoanProductsInterestRecalculationData struct for GetLoanProductsInterestRecalculationData
type GetLoanProductsInterestRecalculationData struct {
	Id                                            int32                                                        `json:"id,omitempty"`
	ProductId                                     int32                                                        `json:"productId,omitempty"`
	InterestRecalculationCompoundingType          GetLoanProductsInterestRecalculationCompoundingType          `json:"interestRecalculationCompoundingType,omitempty"`
	InterestRecalculationCompoundingFrequencyType GetLoanProductsInterestRecalculationCompoundingFrequencyType `json:"interestRecalculationCompoundingFrequencyType,omitempty"`
	RescheduleStrategyType                        GetLoanProductsRescheduleStrategyType                        `json:"rescheduleStrategyType,omitempty"`
	RecalculationRestFrequencyType                GetLoanProductsInterestRecalculationCompoundingFrequencyType `json:"recalculationRestFrequencyType,omitempty"`
	PreClosureInterestCalculationStrategy         GetLoanProductsPreClosureInterestCalculationStrategy         `json:"preClosureInterestCalculationStrategy,omitempty"`
	IsArrearsBasedOnOriginalSchedule              bool                                                         `json:"isArrearsBasedOnOriginalSchedule,omitempty"`
}
