package models

// GetSavingsProductsLiabilityAccountOptions struct for GetSavingsProductsLiabilityAccountOptions
type GetSavingsProductsLiabilityAccountOptions struct {
	Id                   int32                            `json:"id,omitempty"`
	Name                 string                           `json:"name,omitempty"`
	GlCode               int32                            `json:"glCode,omitempty"`
	Disabled             bool                             `json:"disabled,omitempty"`
	ManualEntriesAllowed bool                             `json:"manualEntriesAllowed,omitempty"`
	Type                 GetSavingsProductsLiabilityType  `json:"type,omitempty"`
	Usage                GetSavingsProductsLiabilityUsage `json:"usage,omitempty"`
	NameDecorated        string                           `json:"nameDecorated,omitempty"`
	TagId                GetSavingsProductsLiabilityTagId `json:"tagId,omitempty"`
}
