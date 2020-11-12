package models

// PostSelfLoansLoanIdResponse PostSelfLoansLoanIdResponse
type PostSelfLoansLoanIdResponse struct {
	OfficeId   int32                      `json:"officeId,omitempty"`
	ClientId   int32                      `json:"clientId,omitempty"`
	LoanId     int32                      `json:"loanId,omitempty"`
	ResourceId int32                      `json:"resourceId,omitempty"`
	Changes    PostSelfLoansLoanIdChanges `json:"changes,omitempty"`
}
