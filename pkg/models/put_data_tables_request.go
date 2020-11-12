package models

// PutDataTablesRequest PutDataTablesRequest
type PutDataTablesRequest struct {
	AppTableName  string                              `json:"appTableName,omitempty"`
	DropColumns   []PutDataTablesRequestDropColumns   `json:"dropColumns,omitempty"`
	AddColumns    []PutDataTablesRequestAddColumns    `json:"addColumns,omitempty"`
	ChangeColumns []PutDataTablesRequestChangeColumns `json:"ChangeColumns,omitempty"`
}
