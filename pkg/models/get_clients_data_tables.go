package models

// GetClientsDataTables struct for GetClientsDataTables
type GetClientsDataTables struct {
	ApplicationTableName string                       `json:"applicationTableName,omitempty"`
	RegisteredTableName  string                       `json:"registeredTableName,omitempty"`
	ColumnHeaderData     []GetClientsColumnHeaderData `json:"columnHeaderData,omitempty"`
}
