package models

// GetTellersTellerIdCashiersCashierIdResponse GetTellersTellerIdCashiersCashierIdResponse
type GetTellersTellerIdCashiersCashierIdResponse struct {
	Id          int64  `json:"id,omitempty"`
	TellerId    int64  `json:"tellerId,omitempty"`
	StaffId     int64  `json:"staffId,omitempty"`
	Description string `json:"description,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
	EndDate     string `json:"endDate,omitempty"`
	IsFullDay   bool   `json:"isFullDay,omitempty"`
	StartTime   string `json:"startTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	TellerName  string `json:"tellerName,omitempty"`
	StaffName   string `json:"staffName,omitempty"`
}
