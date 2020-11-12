package models

// GetCentersCenterIdResponse GetCentersCenterIdResponse
type GetCentersCenterIdResponse struct {
	Id         int32            `json:"id,omitempty"`
	Status     GetCentersStatus `json:"status,omitempty"`
	Active     bool             `json:"active,omitempty"`
	Name       string           `json:"name,omitempty"`
	OfficeId   int32            `json:"officeId,omitempty"`
	OfficeName string           `json:"officeName,omitempty"`
	Hierarchy  string           `json:"hierarchy,omitempty"`
}
