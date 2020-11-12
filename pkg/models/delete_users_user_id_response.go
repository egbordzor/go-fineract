package models

// DeleteUsersUserIdResponse DeleteUsersUserIdResponse
type DeleteUsersUserIdResponse struct {
	OfficeId   int64                  `json:"officeId,omitempty"`
	ResourceId int64                  `json:"resourceId,omitempty"`
	Changes    map[string]interface{} `json:"changes,omitempty"`
}
