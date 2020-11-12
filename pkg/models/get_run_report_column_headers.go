package models

// GetRunReportColumnHeaders struct for GetRunReportColumnHeaders
type GetRunReportColumnHeaders struct {
	ColumnName         string `json:"columnName,omitempty"`
	ColumnType         string `json:"columnType,omitempty"`
	IsColumnNullable   bool   `json:"isColumnNullable,omitempty"`
	IsColumnPrimaryKey bool   `json:"isColumnPrimaryKey,omitempty"`
	ColumnValues       string `json:"columnValues,omitempty"`
}
