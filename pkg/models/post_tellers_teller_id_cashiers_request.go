package models

// PostTellersTellerIdCashiersRequest PostTellersTellerIdCashiersRequest
type PostTellersTellerIdCashiersRequest struct {
	EndDate     string `json:"endDate,omitempty"`
	Description string `json:"description,omitempty"`
	IsFullDay   bool   `json:"isFullDay,omitempty"`
	StaffId     int64  `json:"staffId,omitempty"`
	Locale      string `json:"locale,omitempty"`
	DateFormat  string `json:"dateFormat,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
}
