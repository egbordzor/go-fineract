package models

// GetGroupsTemplateColumnHeaderData struct for GetGroupsTemplateColumnHeaderData
type GetGroupsTemplateColumnHeaderData struct {
	ColumnName         string                   `json:"columnName,omitempty"`
	ColumnType         string                   `json:"columnType,omitempty"`
	ColumnLength       int32                    `json:"columnLength,omitempty"`
	ColumnDisplayType  string                   `json:"columnDisplayType,omitempty"`
	IsColumnNullable   bool                     `json:"isColumnNullable,omitempty"`
	IsColumnPrimaryKey bool                     `json:"isColumnPrimaryKey,omitempty"`
	ColumnValues       []map[string]interface{} `json:"columnValues,omitempty"`
}
