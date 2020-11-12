package models

import (
	"time"
)

// JobDetailHistoryDataSwagger struct for JobDetailHistoryDataSwagger
type JobDetailHistoryDataSwagger struct {
	Version         int64     `json:"version,omitempty"`
	JobRunStartTime time.Time `json:"jobRunStartTime,omitempty"`
	JobRunEndTime   time.Time `json:"jobRunEndTime,omitempty"`
	Status          string    `json:"status,omitempty"`
	TriggerType     string    `json:"triggerType,omitempty"`
}
