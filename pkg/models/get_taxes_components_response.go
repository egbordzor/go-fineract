package models

// GetTaxesComponentsResponse GetTaxesComponentsResponse
type GetTaxesComponentsResponse struct {
	Id                     int32                               `json:"id,omitempty"`
	Name                   string                              `json:"name,omitempty"`
	Percentage             float32                             `json:"percentage,omitempty"`
	CreditAccountType      GetTaxesComponentsCreditAccountType `json:"creditAccountType,omitempty"`
	CreditAccount          GetTaxesComponentsCreditAccount     `json:"creditAccount,omitempty"`
	StartDate              string                              `json:"startDate,omitempty"`
	TaxComponentsHistories []map[string]interface{}            `json:"taxComponentsHistories,omitempty"`
}
