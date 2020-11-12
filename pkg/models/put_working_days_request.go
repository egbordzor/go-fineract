package models

// PutWorkingDaysRequest PutWorkingDaysRequest
type PutWorkingDaysRequest struct {
	Recurrence                   string         `json:"recurrence,omitempty"`
	Locale                       string         `json:"locale,omitempty"`
	RepaymentRescheduleType      EnumOptionData `json:"repaymentRescheduleType,omitempty"`
	ExtendTermForDailyRepayments bool           `json:"extendTermForDailyRepayments,omitempty"`
}
