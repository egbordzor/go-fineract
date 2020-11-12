package models

// PutSavingsAccountsChanges struct for PutSavingsAccountsChanges
type PutSavingsAccountsChanges struct {
	NominalAnnualInterestRate float64 `json:"nominalAnnualInterestRate,omitempty"`
	Locale                    string  `json:"locale,omitempty"`
}
