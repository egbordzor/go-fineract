package models

// PutJobsJobIdRequest PutJobsJobsIDRequest
type PutJobsJobIdRequest struct {
	DisplayName    string `json:"displayName,omitempty"`
	CronExpression string `json:"cronExpression,omitempty"`
	Active         bool   `json:"active,omitempty"`
}
