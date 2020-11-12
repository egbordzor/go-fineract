package models

// PutCentersCenterIdResponse PutCentersCenterIdResponse
type PutCentersCenterIdResponse struct {
	OfficeId   int32             `json:"officeId,omitempty"`
	GroupId    int32             `json:"groupId,omitempty"`
	ResourceId int32             `json:"resourceId,omitempty"`
	Changes    PutCentersChanges `json:"changes,omitempty"`
}
