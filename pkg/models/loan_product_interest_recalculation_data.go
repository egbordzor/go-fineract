package models

// LoanProductInterestRecalculationData struct for LoanProductInterestRecalculationData
type LoanProductInterestRecalculationData struct {
	InterestRecalculationCompoundingType      EnumOptionData `json:"interestRecalculationCompoundingType,omitempty"`
	RescheduleStrategyType                    EnumOptionData `json:"rescheduleStrategyType,omitempty"`
	RecalculationRestFrequencyType            EnumOptionData `json:"recalculationRestFrequencyType,omitempty"`
	RecalculationRestFrequencyInterval        int32          `json:"recalculationRestFrequencyInterval,omitempty"`
	RecalculationRestFrequencyNthDay          EnumOptionData `json:"recalculationRestFrequencyNthDay,omitempty"`
	RecalculationRestFrequencyWeekday         EnumOptionData `json:"recalculationRestFrequencyWeekday,omitempty"`
	RecalculationRestFrequencyOnDay           int32          `json:"recalculationRestFrequencyOnDay,omitempty"`
	RecalculationCompoundingFrequencyType     EnumOptionData `json:"recalculationCompoundingFrequencyType,omitempty"`
	RecalculationCompoundingFrequencyInterval int32          `json:"recalculationCompoundingFrequencyInterval,omitempty"`
	RecalculationCompoundingFrequencyNthDay   EnumOptionData `json:"recalculationCompoundingFrequencyNthDay,omitempty"`
	RecalculationCompoundingFrequencyWeekday  EnumOptionData `json:"recalculationCompoundingFrequencyWeekday,omitempty"`
	RecalculationCompoundingFrequencyOnDay    int32          `json:"recalculationCompoundingFrequencyOnDay,omitempty"`
	CompoundingToBePostedAsTransaction        bool           `json:"compoundingToBePostedAsTransaction,omitempty"`
}
