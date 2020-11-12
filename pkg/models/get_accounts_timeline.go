package models

// GetAccountsTimeline struct for GetAccountsTimeline
type GetAccountsTimeline struct {
	SubmittedOnDate      string `json:"submittedOnDate,omitempty"`
	SubmittedByUsername  string `json:"submittedByUsername,omitempty"`
	SubmittedByFirstname string `json:"submittedByFirstname,omitempty"`
	SubmittedByLastname  string `json:"submittedByLastname,omitempty"`
	ApprovedDate         string `json:"approvedDate,omitempty"`
	ApprovedByUsername   string `json:"approvedByUsername,omitempty"`
	ApprovedByFirstname  string `json:"approvedByFirstname,omitempty"`
	ApprovedByLastname   string `json:"approvedByLastname,omitempty"`
	ActivatedDate        string `json:"activatedDate,omitempty"`
}
