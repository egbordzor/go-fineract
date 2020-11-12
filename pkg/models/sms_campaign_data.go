package models

import (
	"time"
)

// SmsCampaignData struct for SmsCampaignData
type SmsCampaignData struct {
	Id                  int64          `json:"id,omitempty"`
	CampaignName        string         `json:"campaignName,omitempty"`
	CampaignType        EnumOptionData `json:"campaignType,omitempty"`
	RunReportId         int64          `json:"runReportId,omitempty"`
	ReportName          string         `json:"reportName,omitempty"`
	ParamValue          string         `json:"paramValue,omitempty"`
	CampaignStatus      EnumOptionData `json:"campaignStatus,omitempty"`
	NextTriggerDate     time.Time      `json:"nextTriggerDate,omitempty"`
	LastTriggerDate     string         `json:"lastTriggerDate,omitempty"`
	RecurrenceStartDate time.Time      `json:"recurrenceStartDate,omitempty"`
	Recurrence          string         `json:"recurrence,omitempty"`
	Notification        bool           `json:"notification,omitempty"`
	Message             string         `json:"message,omitempty"`
}
