package models

// PostCollectionSheetChanges struct for PostCollectionSheetChanges
type PostCollectionSheetChanges struct {
	Locale              string  `json:"locale,omitempty"`
	DateFormat          string  `json:"dateFormat,omitempty"`
	LoanTransactions    []int32 `json:"loanTransactions,omitempty"`
	SavingsTransactions []int32 `json:"SavingsTransactions,omitempty"`
}
