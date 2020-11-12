package models

// PostCodeValuesDataRequest PostCodeValuesDataRequest
type PostCodeValuesDataRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Position    int32  `json:"position,omitempty"`
}
