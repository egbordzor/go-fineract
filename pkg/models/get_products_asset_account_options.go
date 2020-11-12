package models

// GetProductsAssetAccountOptions struct for GetProductsAssetAccountOptions
type GetProductsAssetAccountOptions struct {
	Id                   int32                     `json:"id,omitempty"`
	Name                 string                    `json:"name,omitempty"`
	GlCode               string                    `json:"glCode,omitempty"`
	Disabled             bool                      `json:"disabled,omitempty"`
	ManualEntriesAllowed bool                      `json:"manualEntriesAllowed,omitempty"`
	Type                 GetAssetType              `json:"type,omitempty"`
	Usage                GetProductsLiabilityUsage `json:"usage,omitempty"`
	Description          string                    `json:"description,omitempty"`
	NameDecorated        string                    `json:"nameDecorated,omitempty"`
	TagId                GetProductsTagId          `json:"tagId,omitempty"`
}
