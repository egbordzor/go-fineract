package models

// GetProductsEquityAccountOptions struct for GetProductsEquityAccountOptions
type GetProductsEquityAccountOptions struct {
	Id                   int32                     `json:"id,omitempty"`
	Name                 string                    `json:"name,omitempty"`
	GlCode               string                    `json:"glCode,omitempty"`
	Disabled             bool                      `json:"disabled,omitempty"`
	ManualEntriesAllowed bool                      `json:"manualEntriesAllowed,omitempty"`
	Type                 GetEquityType             `json:"type,omitempty"`
	Usage                GetProductsLiabilityUsage `json:"usage,omitempty"`
	NameDecorated        string                    `json:"nameDecorated,omitempty"`
	TagId                GetProductsTagId          `json:"tagId,omitempty"`
}
