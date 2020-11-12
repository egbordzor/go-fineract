package models

// PostGlAccountsRequest PostGLAccountsRequest
type PostGlAccountsRequest struct {
	Name                 string         `json:"name,omitempty"`
	GlCode               string         `json:"glCode,omitempty"`
	ManualEntriesAllowed bool           `json:"manualEntriesAllowed,omitempty"`
	Type                 string         `json:"type,omitempty"`
	TagId                string         `json:"tagId,omitempty"`
	ParentId             int64          `json:"parentId,omitempty"`
	Usage                EnumOptionData `json:"usage,omitempty"`
	Description          string         `json:"description,omitempty"`
}
