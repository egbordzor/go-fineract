package models

import (
	"time"
)

// PutReportMailingJobsRequest PutReportMailingJobsRequest
type PutReportMailingJobsRequest struct {
	Locale        string    `json:"locale,omitempty"`
	DateFormat    string    `json:"dateFormat,omitempty"`
	StartDateTime time.Time `json:"startDateTime,omitempty"`
}
