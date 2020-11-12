package models

// GetTellersTellerIdCashiersTemplateResponse GetTellersTellerIdCashiersTemplateResponse
type GetTellersTellerIdCashiersTemplateResponse struct {
	TellerId     int64       `json:"tellerId,omitempty"`
	TellerName   string      `json:"tellerName,omitempty"`
	OfficeId     int64       `json:"officeId,omitempty"`
	OfficeName   string      `json:"officeName,omitempty"`
	StaffOptions []StaffData `json:"staffOptions,omitempty"`
}
