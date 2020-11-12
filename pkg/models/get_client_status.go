package models

// GetClientStatus struct for GetClientStatus
type GetClientStatus struct {
	Id          int32  `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}
