package models

// GetWorkingDaysResponse GetWorkingDaysResponse
type GetWorkingDaysResponse struct {
	Id                           int64          `json:"id,omitempty"`
	Recurrence                   string         `json:"recurrence,omitempty"`
	RepaymentRescheduleType      EnumOptionData `json:"repaymentRescheduleType,omitempty"`
	ExtendTermForDailyRepayments bool           `json:"extendTermForDailyRepayments,omitempty"`
}
