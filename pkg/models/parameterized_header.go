package models

// ParameterizedHeader struct for ParameterizedHeader
type ParameterizedHeader struct {
	Value      string            `json:"value,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}
