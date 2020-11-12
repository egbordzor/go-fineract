package models

// PostRepostRequest PostRepostRequest
type PostRepostRequest struct {
	ReportName       string                   `json:"reportName,omitempty"`
	ReportType       string                   `json:"reportType,omitempty"`
	ReportSubType    string                   `json:"reportSubType,omitempty"`
	ReportCategory   string                   `json:"reportCategory,omitempty"`
	Description      string                   `json:"description,omitempty"`
	ReportSql        string                   `json:"reportSql,omitempty"`
	ReportParameters []map[string]interface{} `json:"reportParameters,omitempty"`
}
