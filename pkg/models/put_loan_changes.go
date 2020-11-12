package models

// PutLoanChanges struct for PutLoanChanges
type PutLoanChanges struct {
	Principal float64 `json:"principal,omitempty"`
	Locale    string  `json:"locale,omitempty"`
}
