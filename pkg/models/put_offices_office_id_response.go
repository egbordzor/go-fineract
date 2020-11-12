package models

// PutOfficesOfficeIdResponse PutOfficesOfficeIdResponse
type PutOfficesOfficeIdResponse struct {
	OfficeId   int64                             `json:"officeId,omitempty"`
	ResourceId int64                             `json:"resourceId,omitempty"`
	Changes    PutOfficesOfficeIdResponseChanges `json:"changes,omitempty"`
}
