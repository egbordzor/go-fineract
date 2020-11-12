package models

// PutDataTablesRequestAddColumns struct for PutDataTablesRequestAddColumns
type PutDataTablesRequestAddColumns struct {
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	Code      string `json:"code,omitempty"`
	Mandatory bool   `json:"mandatory,omitempty"`
}
