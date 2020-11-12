package models

// PutTellersResponse PutTellersResponse
type PutTellersResponse struct {
	OfficeId   int64                     `json:"officeId,omitempty"`
	ResourceId int64                     `json:"resourceId,omitempty"`
	Changes    PutTellersResponseChanges `json:"changes,omitempty"`
}
