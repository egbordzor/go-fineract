package models

// PutTaxesComponentsTaxComponentIdRequest PutTaxesComponentsTaxComponentIdRequest
type PutTaxesComponentsTaxComponentIdRequest struct {
	Name       string  `json:"name,omitempty"`
	Percentage float32 `json:"percentage,omitempty"`
	Locale     string  `json:"locale,omitempty"`
	DateFormat string  `json:"dateFormat,omitempty"`
	StartDate  string  `json:"startDate,omitempty"`
}
