package models

// GetTellersTellerIdCashiersResponse GetTellersTellerIdCashiersResponse
type GetTellersTellerIdCashiersResponse struct {
	TellerId   int64         `json:"tellerId,omitempty"`
	TellerName string        `json:"tellerName,omitempty"`
	OfficeId   int64         `json:"officeId,omitempty"`
	OfficeName string        `json:"officeName,omitempty"`
	Cashiers   []CashierData `json:"cashiers,omitempty"`
}
