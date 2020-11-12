package models

// GetCodeValuesDataResponse GetCodeValuesDataResponse
type GetCodeValuesDataResponse struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Position    int32  `json:"position,omitempty"`
}
