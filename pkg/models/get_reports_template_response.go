package models

// GetReportsTemplateResponse GetReportsTemplateResponse
type GetReportsTemplateResponse struct {
	AllowedReportTypes    []string                 `json:"allowedReportTypes,omitempty"`
	AllowedReportSubTypes []string                 `json:"allowedReportSubTypes,omitempty"`
	AllowedParameters     []map[string]interface{} `json:"allowedParameters,omitempty"`
}
