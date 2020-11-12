package models

// PostSelfLoansLoanIdRequest PostSelfLoansLoanIdRequest
type PostSelfLoansLoanIdRequest struct {
	Locale          string `json:"locale,omitempty"`
	DateFormat      string `json:"dateFormat,omitempty"`
	WithdrawnOnDate string `json:"withdrawnOnDate,omitempty"`
	Note            string `json:"note,omitempty"`
}
