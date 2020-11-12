package models

// PutSavingsProductsProductIdRequest PutSavingsProductsProductIdRequest
type PutSavingsProductsProductIdRequest struct {
	Description  string  `json:"description,omitempty"`
	Locale       string  `json:"locale,omitempty"`
	InterestRate float64 `json:"interestRate,omitempty"`
}
