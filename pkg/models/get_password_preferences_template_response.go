package models

// GetPasswordPreferencesTemplateResponse GetPasswordPreferencesTemplateResponse
type GetPasswordPreferencesTemplateResponse struct {
	Id          int64  `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Active      bool   `json:"active,omitempty"`
	Key         string `json:"key,omitempty"`
}
