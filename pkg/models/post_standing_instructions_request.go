package models

// PostStandingInstructionsRequest PostStandingInstructionsRequest
type PostStandingInstructionsRequest struct {
	FromOfficeId        int64  `json:"fromOfficeId,omitempty"`
	FromClientId        int64  `json:"fromClientId,omitempty"`
	FromAccountType     int32  `json:"fromAccountType,omitempty"`
	Name                string `json:"name,omitempty"`
	TransferType        int32  `json:"transferType,omitempty"`
	Priority            int32  `json:"priority,omitempty"`
	Status              int32  `json:"status,omitempty"`
	FromAccountId       int64  `json:"fromAccountId,omitempty"`
	ToOfficeId          int64  `json:"toOfficeId,omitempty"`
	ToClientId          int64  `json:"toClientId,omitempty"`
	ToAccountType       int32  `json:"toAccountType,omitempty"`
	ToAccountId         int64  `json:"toAccountId,omitempty"`
	InstructionType     int32  `json:"instructionType,omitempty"`
	Amount              int32  `json:"amount,omitempty"`
	ValidFrom           string `json:"validFrom,omitempty"`
	RecurrenceType      int32  `json:"recurrenceType,omitempty"`
	RecurrenceInterval  int32  `json:"recurrenceInterval,omitempty"`
	RecurrenceFrequency int32  `json:"recurrenceFrequency,omitempty"`
	// en
	Locale               string `json:"locale,omitempty"`
	DateFormat           string `json:"dateFormat,omitempty"`
	RecurrenceOnMonthDay string `json:"recurrenceOnMonthDay,omitempty"`
	MonthDayFormat       string `json:"monthDayFormat,omitempty"`
}
