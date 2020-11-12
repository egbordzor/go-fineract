package models

// TemplateMapper struct for TemplateMapper
type TemplateMapper struct {
	Id          int64  `json:"id,omitempty"`
	Mapperorder int32  `json:"mapperorder,omitempty"`
	Mapperkey   string `json:"mapperkey,omitempty"`
	Mappervalue string `json:"mappervalue,omitempty"`
	New         bool   `json:"new,omitempty"`
}
