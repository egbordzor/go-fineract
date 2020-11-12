package models

// Permission struct for Permission
type Permission struct {
	Id       int64  `json:"id,omitempty"`
	Grouping string `json:"grouping,omitempty"`
	Code     string `json:"code,omitempty"`
	New      bool   `json:"new,omitempty"`
}
