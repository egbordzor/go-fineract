package models

// PutTaxesGroupTaxGroupIdRequest PutTaxesGroupTaxGroupIdRequest
type PutTaxesGroupTaxGroupIdRequest struct {
	Name          string                       `json:"name,omitempty"`
	Locale        string                       `json:"locale,omitempty"`
	TaxComponents []PutTaxesGroupTaxComponents `json:"taxComponents,omitempty"`
	DateFormat    string                       `json:"dateFormat,omitempty"`
}
