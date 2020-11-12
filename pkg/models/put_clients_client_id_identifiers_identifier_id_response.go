package models

// PutClientsClientIdIdentifiersIdentifierIdResponse PutClientsClientIdIdentifiersIdentifierIdResponse
type PutClientsClientIdIdentifiersIdentifierIdResponse struct {
	OfficeId   int32                                            `json:"officeId,omitempty"`
	ClientId   int32                                            `json:"clientId,omitempty"`
	ResourceId int32                                            `json:"resourceId,omitempty"`
	Changes    PutClientsClientIdIdentifiersIdentifierIdRequest `json:"changes,omitempty"`
}
