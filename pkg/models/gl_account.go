package models

// GlAccount struct for GlAccount
type GlAccount struct {
	Id                   int64       `json:"id,omitempty"`
	Children             []GlAccount `json:"children,omitempty"`
	Name                 string      `json:"name,omitempty"`
	GlCode               string      `json:"glCode,omitempty"`
	Disabled             bool        `json:"disabled,omitempty"`
	ManualEntriesAllowed bool        `json:"manualEntriesAllowed,omitempty"`
	Type                 int32       `json:"type,omitempty"`
	Usage                int32       `json:"usage,omitempty"`
	HeaderAccount        bool        `json:"headerAccount,omitempty"`
	DetailAccount        bool        `json:"detailAccount,omitempty"`
	New                  bool        `json:"new,omitempty"`
}
