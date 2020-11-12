package models

// PostLoansLoanIdScheduleResponse PostLoansLoanIdScheduleResponse
type PostLoansLoanIdScheduleResponse struct {
	LoanId  int32           `json:"loanId,omitempty"`
	Changes PostLoanChanges `json:"changes,omitempty"`
}
