package models

import (
	"time"
)

// GetReportMailingJobsResponse GetReportMailingJobsResponse
type GetReportMailingJobsResponse struct {
	Id                        int64                  `json:"id,omitempty"`
	Name                      string                 `json:"name,omitempty"`
	Description               string                 `json:"description,omitempty"`
	StartDateTime             time.Time              `json:"startDateTime,omitempty"`
	Recurrence                string                 `json:"recurrence,omitempty"`
	Timeline                  map[string]interface{} `json:"timeline,omitempty"`
	EmailRecipients           string                 `json:"emailRecipients,omitempty"`
	EmailSubject              string                 `json:"emailSubject,omitempty"`
	EmailMessage              string                 `json:"emailMessage,omitempty"`
	EmailAttachmentFileFormat EnumOptionData         `json:"emailAttachmentFileFormat,omitempty"`
	StretchyReport            map[string]interface{} `json:"stretchyReport,omitempty"`
	StretchyReportParamMap    string                 `json:"stretchyReportParamMap,omitempty"`
	NextRunDateTime           time.Time              `json:"nextRunDateTime,omitempty"`
	NumberOfRuns              int32                  `json:"numberOfRuns,omitempty"`
	IsActive                  bool                   `json:"isActive,omitempty"`
	RunAsUserId               int64                  `json:"runAsUserId,omitempty"`
}
