package models

// GlAccountData struct for GlAccountData
type GlAccountData struct {
	Id       int64          `json:"id,omitempty"`
	Name     string         `json:"name,omitempty"`
	GlCode   string         `json:"glCode,omitempty"`
	Type     EnumOptionData `json:"type,omitempty"`
	TagId    CodeValueData  `json:"tagId,omitempty"`
	RowIndex int32          `json:"rowIndex,omitempty"`
	TypeId   int32          `json:"typeId,omitempty"`
}
