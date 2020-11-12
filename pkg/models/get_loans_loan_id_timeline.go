package models

// GetLoansLoanIdTimeline struct for GetLoansLoanIdTimeline
type GetLoansLoanIdTimeline struct {
	SubmittedOnDate          string `json:"submittedOnDate,omitempty"`
	SubmittedByUsername      string `json:"submittedByUsername,omitempty"`
	SubmittedByFirstname     string `json:"submittedByFirstname,omitempty"`
	SubmittedByLastname      string `json:"submittedByLastname,omitempty"`
	ApprovedOnDate           string `json:"approvedOnDate,omitempty"`
	ApprovedByUsername       string `json:"approvedByUsername,omitempty"`
	ApprovedByFirstname      string `json:"approvedByFirstname,omitempty"`
	ApprovedByLastname       string `json:"approvedByLastname,omitempty"`
	ExpectedDisbursementDate string `json:"expectedDisbursementDate,omitempty"`
	ActualDisbursementDate   string `json:"actualDisbursementDate,omitempty"`
	DisbursedByUsername      string `json:"disbursedByUsername,omitempty"`
	DisbursedByFirstname     string `json:"disbursedByFirstname,omitempty"`
	DisbursedByLastname      string `json:"disbursedByLastname,omitempty"`
	ExpectedMaturityDate     string `json:"expectedMaturityDate,omitempty"`
}
