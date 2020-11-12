package models

// PutTaxesComponentsChanges struct for PutTaxesComponentsChanges
type PutTaxesComponentsChanges struct {
	Percentage float32 `json:"percentage,omitempty"`
	Name       string  `json:"name,omitempty"`
	StartDate  string  `json:"startDate,omitempty"`
}
