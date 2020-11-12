package models

// PutSelfLoansChanges struct for PutSelfLoansChanges
type PutSelfLoansChanges struct {
	Principal int64  `json:"principal,omitempty"`
	Locale    string `json:"locale,omitempty"`
}
