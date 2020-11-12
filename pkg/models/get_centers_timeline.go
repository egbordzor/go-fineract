package models

// GetCentersTimeline struct for GetCentersTimeline
type GetCentersTimeline struct {
	SubmittedOnDate      string `json:"submittedOnDate,omitempty"`
	SubmittedByUsername  string `json:"submittedByUsername,omitempty"`
	SubmittedByFirstname string `json:"submittedByFirstname,omitempty"`
	SubmittedByLastname  string `json:"submittedByLastname,omitempty"`
}
