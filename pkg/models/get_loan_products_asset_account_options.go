package models

// GetLoanProductsAssetAccountOptions struct for GetLoanProductsAssetAccountOptions
type GetLoanProductsAssetAccountOptions struct {
	Id                         int32                         `json:"id,omitempty"`
	Name                       string                        `json:"name,omitempty"`
	GlCode                     int32                         `json:"glCode,omitempty"`
	Disabled                   bool                          `json:"disabled,omitempty"`
	ManualEntriesAllowed       bool                          `json:"manualEntriesAllowed,omitempty"`
	Type                       GetLoanProductsLiabilityType  `json:"type,omitempty"`
	Usage                      GetLoanProductsLiabilityUsage `json:"usage,omitempty"`
	NameDecorated              string                        `json:"nameDecorated,omitempty"`
	TagId                      GetLoanProductsLiabilityTagId `json:"tagId,omitempty"`
	OrganizationRunningBalance int32                         `json:"organizationRunningBalance,omitempty"`
}
