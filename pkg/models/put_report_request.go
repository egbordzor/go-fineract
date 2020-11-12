package models

// PutReportRequest PutReportRequest
type PutReportRequest struct {
	ReportName       string                   `json:"reportName,omitempty"`
	ReportParameters []map[string]interface{} `json:"reportParameters,omitempty"`
}
