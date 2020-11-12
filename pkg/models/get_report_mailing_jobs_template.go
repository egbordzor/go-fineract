package models

// GetReportMailingJobsTemplate GetReportMailingJobsTemplate
type GetReportMailingJobsTemplate struct {
	IsActive                         bool             `json:"isActive,omitempty"`
	EmailAttachmentFileFormatOptions []EnumOptionData `json:"emailAttachmentFileFormatOptions,omitempty"`
	StretchyReportParamDateOptions   []EnumOptionData `json:"stretchyReportParamDateOptions,omitempty"`
}
