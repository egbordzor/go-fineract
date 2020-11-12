package models

// PutGroupsGroupIdResponse PutGroupsGroupIdResponse
type PutGroupsGroupIdResponse struct {
	OfficeId   int32                   `json:"officeId,omitempty"`
	GroupId    int32                   `json:"groupId,omitempty"`
	ResourceId int32                   `json:"resourceId,omitempty"`
	Changes    PutGroupsGroupIdChanges `json:"changes,omitempty"`
}
