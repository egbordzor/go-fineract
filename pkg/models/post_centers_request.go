package models

// PostCentersRequest PostCentersRequest
type PostCentersRequest struct {
	Name     string `json:"name,omitempty"`
	OfficeId int32  `json:"officeId,omitempty"`
	Active   bool   `json:"active,omitempty"`
}
