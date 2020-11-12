package models

// GetCodesResponse GetCodesResponse
type GetCodesResponse struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	SystemDefined bool   `json:"systemDefined,omitempty"`
}
