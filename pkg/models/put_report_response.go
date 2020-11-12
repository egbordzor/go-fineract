package models

// PutReportResponse PutReportResponse
type PutReportResponse struct {
	ResourceId int64                    `json:"resourceId,omitempty"`
	Changes    PutReportResponseChanges `json:"changes,omitempty"`
}
