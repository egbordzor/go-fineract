package models

// PostGroupsGroupIdCommandUnassignStaffResponse PostGroupsGroupIdCommandUnassignStaffResponse
type PostGroupsGroupIdCommandUnassignStaffResponse struct {
	OfficeId   int32                  `json:"officeId,omitempty"`
	GroupId    int32                  `json:"groupId,omitempty"`
	ResourceId int32                  `json:"resourceId,omitempty"`
	Changes    map[string]interface{} `json:"changes,omitempty"`
}
