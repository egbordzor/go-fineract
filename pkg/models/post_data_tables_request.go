package models

// PostDataTablesRequest PostDataTablesRequest
type PostDataTablesRequest struct {
	ApplicationTableName string `json:"applicationTableName,omitempty"`
	RegisteredTableName  string `json:"registeredTableName,omitempty"`
	// Allows to create multiple entries in the Data Table. Optional, defaults to false. If this property is not provided Data Table will allow only one entry.
	MultiRow         bool                        `json:"multiRow,omitempty"`
	ColumnHeaderData []ResultsetColumnHeaderData `json:"columnHeaderData,omitempty"`
}
