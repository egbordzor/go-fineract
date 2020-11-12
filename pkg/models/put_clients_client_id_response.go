package models

// PutClientsClientIdResponse PutClientsClientIdResponse
type PutClientsClientIdResponse struct {
	OfficeId   int32                     `json:"officeId,omitempty"`
	ClientId   int32                     `json:"clientId,omitempty"`
	ResourceId int32                     `json:"resourceId,omitempty"`
	Changes    PutClientsClientIdRequest `json:"changes,omitempty"`
}
