package models

// PutReportResponseChanges struct for PutReportResponseChanges
type PutReportResponseChanges struct {
	ReportName       string                   `json:"reportName,omitempty"`
	ReportParameters []map[string]interface{} `json:"reportParameters,omitempty"`
}
