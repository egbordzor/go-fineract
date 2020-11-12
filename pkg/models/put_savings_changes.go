package models

// PutSavingsChanges struct for PutSavingsChanges
type PutSavingsChanges struct {
	Description  string  `json:"description,omitempty"`
	InterestRate float64 `json:"interestRate,omitempty"`
	Locale       string  `json:"locale,omitempty"`
}
