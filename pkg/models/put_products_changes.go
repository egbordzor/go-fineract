package models

// PutProductsChanges struct for PutProductsChanges
type PutProductsChanges struct {
	Description string  `json:"description,omitempty"`
	UnitPrice   float64 `json:"unitPrice,omitempty"`
	Locale      string  `json:"locale,omitempty"`
}
