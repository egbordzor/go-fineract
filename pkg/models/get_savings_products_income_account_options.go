package models

// GetSavingsProductsIncomeAccountOptions struct for GetSavingsProductsIncomeAccountOptions
type GetSavingsProductsIncomeAccountOptions struct {
	Id                   int32                            `json:"id,omitempty"`
	Name                 string                           `json:"name,omitempty"`
	GlCode               int32                            `json:"glCode,omitempty"`
	Disabled             bool                             `json:"disabled,omitempty"`
	ManualEntriesAllowed bool                             `json:"manualEntriesAllowed,omitempty"`
	Type                 GetSavingsProductsIncomeType     `json:"type,omitempty"`
	Usage                GetSavingsProductsLiabilityUsage `json:"usage,omitempty"`
	TagId                map[string]interface{}           `json:"tagId,omitempty"`
}
