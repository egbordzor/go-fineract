package models

// PutTaxesGroupChanges struct for PutTaxesGroupChanges
type PutTaxesGroupChanges struct {
	AddComponents      []int32                           `json:"addComponents,omitempty"`
	ModifiedComponents []PutTaxesGroupModifiedComponents `json:"modifiedComponents,omitempty"`
	Name               string                            `json:"name,omitempty"`
}
