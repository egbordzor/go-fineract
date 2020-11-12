package models

// PostLoansLoanIdRequest PostLoansLoanIdRequest
type PostLoansLoanIdRequest struct {
	ToLoanOfficerId   int32  `json:"toLoanOfficerId,omitempty"`
	AssignmentDate    string `json:"assignmentDate,omitempty"`
	Locale            string `json:"locale,omitempty"`
	DateFormat        string `json:"dateFormat,omitempty"`
	FromLoanOfficerId int32  `json:"fromLoanOfficerId,omitempty"`
}
