package models

// GetClientsTimeline struct for GetClientsTimeline
type GetClientsTimeline struct {
	SubmittedOnDate      string `json:"submittedOnDate,omitempty"`
	SubmittedByUsername  string `json:"submittedByUsername,omitempty"`
	SubmittedByFirstname string `json:"submittedByFirstname,omitempty"`
	SubmittedByLastname  string `json:"submittedByLastname,omitempty"`
	ActivatedOnDate      string `json:"activatedOnDate,omitempty"`
	ActivatedByUsername  string `json:"activatedByUsername,omitempty"`
	ActivatedByFirstname string `json:"activatedByFirstname,omitempty"`
	ActivatedByLastname  string `json:"activatedByLastname,omitempty"`
}
