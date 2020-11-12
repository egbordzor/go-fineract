package models

// GetAccountsLinkedToPocketResponse GetAccountsLinkedToPocketResponse
type GetAccountsLinkedToPocketResponse struct {
	LoanAccounts   []GetPocketLoanAccounts   `json:"loanAccounts,omitempty"`
	SavingAccounts []GetPocketSavingAccounts `json:"savingAccounts,omitempty"`
	ShareAccounts  []map[string]interface{}  `json:"shareAccounts,omitempty"`
}
