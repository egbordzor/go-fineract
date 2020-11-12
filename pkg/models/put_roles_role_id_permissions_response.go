package models

// PutRolesRoleIdPermissionsResponse PutRolesRoleIdPermissionsResponse
type PutRolesRoleIdPermissionsResponse struct {
	ResourceId  int64                                         `json:"resourceId,omitempty"`
	Permissions PostRolesRoleIdPermissionsResponsePermissions `json:"permissions,omitempty"`
}
