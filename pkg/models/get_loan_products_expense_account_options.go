package models

// GetLoanProductsExpenseAccountOptions struct for GetLoanProductsExpenseAccountOptions
type GetLoanProductsExpenseAccountOptions struct {
	Id                         int32                         `json:"id,omitempty"`
	Name                       string                        `json:"name,omitempty"`
	GlCode                     int32                         `json:"glCode,omitempty"`
	Disabled                   bool                          `json:"disabled,omitempty"`
	ManualEntriesAllowed       bool                          `json:"manualEntriesAllowed,omitempty"`
	Type                       GetLoanProductsExpenseType    `json:"type,omitempty"`
	Usage                      GetLoanProductsLiabilityUsage `json:"usage,omitempty"`
	NameDecorated              string                        `json:"nameDecorated,omitempty"`
	TagId                      GetLoanProductsLiabilityTagId `json:"tagId,omitempty"`
	OrganizationRunningBalance int32                         `json:"organizationRunningBalance,omitempty"`
}
