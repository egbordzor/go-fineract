package models

// PutTaxesGroupTaxComponents struct for PutTaxesGroupTaxComponents
type PutTaxesGroupTaxComponents struct {
	Id             int32  `json:"id,omitempty"`
	TaxComponentId int32  `json:"taxComponentId,omitempty"`
	EndDate        string `json:"endDate,omitempty"`
}
