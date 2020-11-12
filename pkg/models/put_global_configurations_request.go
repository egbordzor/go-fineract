package models

// PutGlobalConfigurationsRequest PutGlobalConfigurationsRequest
type PutGlobalConfigurationsRequest struct {
	Enabled     bool  `json:"enabled,omitempty"`
	Description int64 `json:"description,omitempty"`
}
