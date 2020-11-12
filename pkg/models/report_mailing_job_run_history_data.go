package models

import (
	"time"
)

// ReportMailingJobRunHistoryData struct for ReportMailingJobRunHistoryData
type ReportMailingJobRunHistoryData struct {
	Id                 int64     `json:"id,omitempty"`
	ReportMailingJobId int64     `json:"reportMailingJobId,omitempty"`
	StartDateTime      time.Time `json:"startDateTime,omitempty"`
	EndDateTime        time.Time `json:"endDateTime,omitempty"`
	Status             string    `json:"status,omitempty"`
	ErrorMessage       string    `json:"errorMessage,omitempty"`
	ErrorLog           string    `json:"errorLog,omitempty"`
}
