package models

// GetReportNameResponse GetReportNameResponse
type GetReportNameResponse struct {
	ColumnHeaders ResultsetColumnHeaderData `json:"columnHeaders,omitempty"`
	Row           ResultsetRowData          `json:"row,omitempty"`
}
