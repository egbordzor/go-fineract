package models

import (
	"time"
)

// PostReportMailingJobsRequest PostReportMailingJobsRequest
type PostReportMailingJobsRequest struct {
	Locale                 string    `json:"locale,omitempty"`
	DateFormat             string    `json:"dateFormat,omitempty"`
	Name                   string    `json:"name,omitempty"`
	Description            string    `json:"description,omitempty"`
	StartDateTime          time.Time `json:"startDateTime,omitempty"`
	StretchyReportId       int64     `json:"stretchyReportId,omitempty"`
	EmailRecipients        string    `json:"emailRecipients,omitempty"`
	EmailSubject           string    `json:"emailSubject,omitempty"`
	EmailMessage           string    `json:"emailMessage,omitempty"`
	Recurrence             string    `json:"recurrence,omitempty"`
	IsActive               bool      `json:"isActive,omitempty"`
	StretchyReportParamMap string    `json:"stretchyReportParamMap,omitempty"`
}
