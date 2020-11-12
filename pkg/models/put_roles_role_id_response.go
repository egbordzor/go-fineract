package models

// PutRolesRoleIdResponse PutRolesRoleIdResponse
type PutRolesRoleIdResponse struct {
	ResourceId int64                         `json:"resourceId,omitempty"`
	Changes    PutRolesRoleIdResponseChanges `json:"changes,omitempty"`
}
