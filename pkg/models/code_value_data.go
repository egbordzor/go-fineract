package models

// CodeValueData struct for CodeValueData
type CodeValueData struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Active    bool   `json:"active,omitempty"`
	Mandatory bool   `json:"mandatory,omitempty"`
}
