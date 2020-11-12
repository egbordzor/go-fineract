package models

// GetDataTablesAppTableIdResponse GetDataTablesAppTableIdResponse
type GetDataTablesAppTableIdResponse struct {
	ColumnHeaders []ResultsetColumnHeaderData `json:"columnHeaders,omitempty"`
	Data          []ResultsetRowData          `json:"data,omitempty"`
}
