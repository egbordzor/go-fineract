package models

// GetLoanProductsIncomeAccountOptions struct for GetLoanProductsIncomeAccountOptions
type GetLoanProductsIncomeAccountOptions struct {
	Id                         int32                         `json:"id,omitempty"`
	Name                       string                        `json:"name,omitempty"`
	GlCode                     int32                         `json:"glCode,omitempty"`
	Disabled                   bool                          `json:"disabled,omitempty"`
	ManualEntriesAllowed       bool                          `json:"manualEntriesAllowed,omitempty"`
	Type                       GetLoanProductsIncomeType     `json:"type,omitempty"`
	Usage                      GetLoanProductsLiabilityUsage `json:"usage,omitempty"`
	NameDecorated              string                        `json:"nameDecorated,omitempty"`
	TagId                      GetLoanProductsLiabilityTagId `json:"tagId,omitempty"`
	OrganizationRunningBalance int32                         `json:"organizationRunningBalance,omitempty"`
}
