package models

// PutUsersUserIdResponse PutUsersUserIdResponse
type PutUsersUserIdResponse struct {
	OfficeId   int64                         `json:"officeId,omitempty"`
	ResourceId int64                         `json:"resourceId,omitempty"`
	Changes    PutUsersUserIdResponseChanges `json:"changes,omitempty"`
}
