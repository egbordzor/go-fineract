package models

// GetClientsColumnHeaderData struct for GetClientsColumnHeaderData
type GetClientsColumnHeaderData struct {
	ColumnName         string   `json:"columnName,omitempty"`
	ColumnType         string   `json:"columnType,omitempty"`
	ColumnLength       int32    `json:"columnLength,omitempty"`
	ColumnDisplayType  string   `json:"columnDisplayType,omitempty"`
	IsColumnNullable   bool     `json:"isColumnNullable,omitempty"`
	IsColumnPrimaryKey bool     `json:"isColumnPrimaryKey,omitempty"`
	ColumnValues       []string `json:"columnValues,omitempty"`
}
