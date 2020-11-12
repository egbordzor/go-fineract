package models

// GetJobsJobIdJobRunHistoryResponse GetJobsJobIDJobRunHistoryResponse
type GetJobsJobIdJobRunHistoryResponse struct {
	TotalFilteredRecords int32                         `json:"totalFilteredRecords,omitempty"`
	PageItems            []JobDetailHistoryDataSwagger `json:"pageItems,omitempty"`
}
