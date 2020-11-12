package models

import (
	"time"
)

// GetJobsResponse GetJobsResponse
type GetJobsResponse struct {
	JobId             int64                  `json:"jobId,omitempty"`
	DisplayName       string                 `json:"displayName,omitempty"`
	NextRunTime       time.Time              `json:"nextRunTime,omitempty"`
	InitializingError string                 `json:"initializingError,omitempty"`
	CronExpression    string                 `json:"cronExpression,omitempty"`
	Active            bool                   `json:"active,omitempty"`
	CurrentlyRunning  bool                   `json:"currentlyRunning,omitempty"`
	LastRunHistory    map[string]interface{} `json:"lastRunHistory,omitempty"`
}
