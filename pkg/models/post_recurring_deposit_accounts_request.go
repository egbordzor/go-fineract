package models

// PostRecurringDepositAccountsRequest PostRecurringDepositAccountsRequest
type PostRecurringDepositAccountsRequest struct {
	ClientId                          int32   `json:"clientId,omitempty"`
	ProductId                         int32   `json:"productId,omitempty"`
	Locale                            string  `json:"locale,omitempty"`
	DateFormat                        string  `json:"dateFormat,omitempty"`
	SubmittedOnDate                   string  `json:"submittedOnDate,omitempty"`
	DepositPeriod                     int32   `json:"depositPeriod,omitempty"`
	DepositPeriodFrequencyId          int32   `json:"depositPeriodFrequencyId,omitempty"`
	DepositAmount                     float32 `json:"depositAmount,omitempty"`
	IsCalendarInherited               bool    `json:"isCalendarInherited,omitempty"`
	RecurringFrequency                int32   `json:"recurringFrequency,omitempty"`
	RecurringFrequencyType            int32   `json:"recurringFrequencyType,omitempty"`
	MandatoryRecommendedDepositAmount int64   `json:"mandatoryRecommendedDepositAmount,omitempty"`
}
