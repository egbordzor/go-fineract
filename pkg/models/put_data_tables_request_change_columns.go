package models

// PutDataTablesRequestChangeColumns struct for PutDataTablesRequestChangeColumns
type PutDataTablesRequestChangeColumns struct {
	Name      string `json:"name,omitempty"`
	NewName   string `json:"newName,omitempty"`
	Code      string `json:"code,omitempty"`
	NewCode   string `json:"newCode,omitempty"`
	Mandatory bool   `json:"mandatory,omitempty"`
}
