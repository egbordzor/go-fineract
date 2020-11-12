package models

// PutProductsTypeProductIdRequest PutProductsTypeProductIdRequest
type PutProductsTypeProductIdRequest struct {
	Description string  `json:"description,omitempty"`
	Locale      string  `json:"locale,omitempty"`
	UnitPrice   float64 `json:"unitPrice,omitempty"`
}
