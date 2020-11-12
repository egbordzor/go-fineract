package models

// ChargeFeeOnMonthDay struct for ChargeFeeOnMonthDay
type ChargeFeeOnMonthDay struct {
	Month      string `json:"month,omitempty"`
	MonthValue int32  `json:"monthValue,omitempty"`
	DayOfMonth int32  `json:"dayOfMonth,omitempty"`
}
