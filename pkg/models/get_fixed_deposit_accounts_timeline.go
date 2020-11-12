package models

// GetFixedDepositAccountsTimeline struct for GetFixedDepositAccountsTimeline
type GetFixedDepositAccountsTimeline struct {
	SubmittedOnDate      string `json:"submittedOnDate,omitempty"`
	SubmittedByUsername  string `json:"submittedByUsername,omitempty"`
	SubmittedByFirstname string `json:"submittedByFirstname,omitempty"`
	SubmittedByLastname  string `json:"submittedByLastname,omitempty"`
}
