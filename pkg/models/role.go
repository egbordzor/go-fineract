package models

// Role struct for Role
type Role struct {
	Id          int64        `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Disabled    bool         `json:"disabled,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
	Enabled     bool         `json:"enabled,omitempty"`
	New         bool         `json:"new,omitempty"`
}
