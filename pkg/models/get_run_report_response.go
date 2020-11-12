package models

// GetRunReportResponse GetRunReportResponse
type GetRunReportResponse struct {
	ColumnHeaders []GetRunReportColumnHeaders `json:"columnHeaders,omitempty"`
	Data          []GetPocketData             `json:"data,omitempty"`
}
