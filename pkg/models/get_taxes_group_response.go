package models

// GetTaxesGroupResponse GetTaxesGroupResponse
type GetTaxesGroupResponse struct {
	Id              int32                          `json:"id,omitempty"`
	Name            string                         `json:"name,omitempty"`
	TaxAssociations []GetTaxesGroupTaxAssociations `json:"taxAssociations,omitempty"`
}
