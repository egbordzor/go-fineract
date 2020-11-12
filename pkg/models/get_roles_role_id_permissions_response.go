package models

// GetRolesRoleIdPermissionsResponse GetRolesRoleIdPermissionsResponse
type GetRolesRoleIdPermissionsResponse struct {
	Id                  int64                                             `json:"id,omitempty"`
	Name                string                                            `json:"name,omitempty"`
	Description         string                                            `json:"description,omitempty"`
	PermissionUsageData []GetRolesRoleIdPermissionsResponsePermissionData `json:"permissionUsageData,omitempty"`
}
