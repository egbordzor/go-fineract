package models

// PutCodeValuesDataRequest PutCodeValuesDataRequest
type PutCodeValuesDataRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Position    int32  `json:"position,omitempty"`
}
