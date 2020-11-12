package models

// PutReportMailingJobsResponse PutReportMailingJobsResponse
type PutReportMailingJobsResponse struct {
	ResourceId int64                               `json:"resourceId,omitempty"`
	Changes    PutReportMailingJobsResponseChanges `json:"changes,omitempty"`
}
