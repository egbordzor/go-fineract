package models

// GetReportsResponse GetReportsResponse
type GetReportsResponse struct {
	Id               int64                    `json:"id,omitempty"`
	ReportName       string                   `json:"reportName,omitempty"`
	ReportType       string                   `json:"reportType,omitempty"`
	ReportSubType    string                   `json:"reportSubType,omitempty"`
	ReportCategory   string                   `json:"reportCategory,omitempty"`
	Description      string                   `json:"description,omitempty"`
	ReportSql        string                   `json:"reportSql,omitempty"`
	CoreReport       bool                     `json:"coreReport,omitempty"`
	UseReport        bool                     `json:"useReport,omitempty"`
	ReportParameters []map[string]interface{} `json:"reportParameters,omitempty"`
}
