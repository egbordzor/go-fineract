package models

// PostTaxesGroupRequest PostTaxesGroupRequest
type PostTaxesGroupRequest struct {
	Name          string                        `json:"name,omitempty"`
	Locale        string                        `json:"locale,omitempty"`
	TaxComponents []PostTaxesGroupTaxComponents `json:"taxComponents,omitempty"`
	DateFormat    string                        `json:"dateFormat,omitempty"`
}
