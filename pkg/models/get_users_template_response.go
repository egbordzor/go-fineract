package models

// GetUsersTemplateResponse GetUsersTemplateResponse
type GetUsersTemplateResponse struct {
	AllowedOffices   []OfficeData `json:"allowedOffices,omitempty"`
	AvailableRoles   []RoleData   `json:"availableRoles,omitempty"`
	SelfServiceRoles []RoleData   `json:"selfServiceRoles,omitempty"`
}
