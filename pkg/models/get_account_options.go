package models

// GetAccountOptions struct for GetAccountOptions
type GetAccountOptions struct {
	Id          int32  `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}
