package models

// GetGlAccountsResponse GetGLAccountsResponse
type GetGlAccountsResponse struct {
	Id                         int64          `json:"id,omitempty"`
	Name                       string         `json:"name,omitempty"`
	ParentId                   int64          `json:"parentId,omitempty"`
	GlCode                     string         `json:"glCode,omitempty"`
	Disabled                   bool           `json:"disabled,omitempty"`
	ManualEntriesAllowed       bool           `json:"manualEntriesAllowed,omitempty"`
	Type                       EnumOptionData `json:"type,omitempty"`
	Usage                      EnumOptionData `json:"usage,omitempty"`
	Description                string         `json:"description,omitempty"`
	NameDecorated              string         `json:"nameDecorated,omitempty"`
	TagId                      CodeValueData  `json:"tagId,omitempty"`
	OrganizationRunningBalance int64          `json:"organizationRunningBalance,omitempty"`
}
