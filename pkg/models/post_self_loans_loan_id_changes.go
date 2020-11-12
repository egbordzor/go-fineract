package models

// PostSelfLoansLoanIdChanges struct for PostSelfLoansLoanIdChanges
type PostSelfLoansLoanIdChanges struct {
	Status          PostSelfLoansLoanIdStatus `json:"status,omitempty"`
	Locale          string                    `json:"locale,omitempty"`
	DateFormat      string                    `json:"dateFormat,omitempty"`
	WithdrawnOnDate string                    `json:"withdrawnOnDate,omitempty"`
	ClosedOnDate    string                    `json:"closedOnDate,omitempty"`
}
