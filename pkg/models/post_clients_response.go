package models

// PostClientsResponse PostClientsResponse
type PostClientsResponse struct {
	OfficeId   int32 `json:"officeId,omitempty"`
	GroupId    int32 `json:"groupId,omitempty"`
	ClientId   int32 `json:"clientId,omitempty"`
	ResourceId int32 `json:"resourceId,omitempty"`
}
