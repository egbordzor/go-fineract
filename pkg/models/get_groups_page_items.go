package models

// GetGroupsPageItems struct for GetGroupsPageItems
type GetGroupsPageItems struct {
	Id         int32           `json:"id,omitempty"`
	Name       string          `json:"name,omitempty"`
	Status     GetGroupsStatus `json:"status,omitempty"`
	Active     bool            `json:"active,omitempty"`
	OfficeId   int32           `json:"officeId,omitempty"`
	OfficeName string          `json:"officeName,omitempty"`
	Hierarchy  string          `json:"hierarchy,omitempty"`
}
