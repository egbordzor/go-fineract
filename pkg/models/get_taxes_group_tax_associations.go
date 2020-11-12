package models

// GetTaxesGroupTaxAssociations struct for GetTaxesGroupTaxAssociations
type GetTaxesGroupTaxAssociations struct {
	Id           int32                     `json:"id,omitempty"`
	TaxComponent GetTaxesGroupTaxComponent `json:"taxComponent,omitempty"`
	StartDate    string                    `json:"startDate,omitempty"`
}
