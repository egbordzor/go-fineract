package models

// GetRolesRoleIdPermissionsResponsePermissionData struct for GetRolesRoleIdPermissionsResponsePermissionData
type GetRolesRoleIdPermissionsResponsePermissionData struct {
	Grouping   string `json:"grouping,omitempty"`
	Code       string `json:"code,omitempty"`
	EntityName string `json:"entityName,omitempty"`
	ActionName string `json:"actionName,omitempty"`
	Selected   bool   `json:"selected,omitempty"`
}
