package models

// PostGroupsRequest PostGroupsRequest
type PostGroupsRequest struct {
	OfficeId int32  `json:"officeId,omitempty"`
	Name     string `json:"name,omitempty"`
	Active   bool   `json:"active,omitempty"`
}
