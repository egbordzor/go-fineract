package models

// LoanProductInterestRecalculationDetails struct for LoanProductInterestRecalculationDetails
type LoanProductInterestRecalculationDetails struct {
	Id                                     int64  `json:"id,omitempty"`
	InterestRecalculationCompoundingMethod int32  `json:"interestRecalculationCompoundingMethod,omitempty"`
	RescheduleStrategyMethod               int32  `json:"rescheduleStrategyMethod,omitempty"`
	RestFrequencyType                      string `json:"restFrequencyType,omitempty"`
	RestInterval                           int32  `json:"restInterval,omitempty"`
	RestFrequencyNthDay                    int32  `json:"restFrequencyNthDay,omitempty"`
	RestFrequencyWeekday                   int32  `json:"restFrequencyWeekday,omitempty"`
	RestFrequencyOnDay                     int32  `json:"restFrequencyOnDay,omitempty"`
	CompoundingFrequencyType               string `json:"compoundingFrequencyType,omitempty"`
	CompoundingInterval                    int32  `json:"compoundingInterval,omitempty"`
	CompoundingFrequencyNthDay             int32  `json:"compoundingFrequencyNthDay,omitempty"`
	CompoundingFrequencyWeekday            int32  `json:"compoundingFrequencyWeekday,omitempty"`
	CompoundingFrequencyOnDay              int32  `json:"compoundingFrequencyOnDay,omitempty"`
	IsCompoundingToBePostedAsTransaction   bool   `json:"isCompoundingToBePostedAsTransaction,omitempty"`
	ArrearsBasedOnOriginalSchedule         bool   `json:"arrearsBasedOnOriginalSchedule,omitempty"`
	New                                    bool   `json:"new,omitempty"`
}
