package models

// GetRecurringDepositAccountsTimeline struct for GetRecurringDepositAccountsTimeline
type GetRecurringDepositAccountsTimeline struct {
	SubmittedOnDate      string `json:"submittedOnDate,omitempty"`
	SubmittedByUsername  string `json:"submittedByUsername,omitempty"`
	SubmittedByFirstname string `json:"submittedByFirstname,omitempty"`
	SubmittedByLastname  string `json:"submittedByLastname,omitempty"`
}
