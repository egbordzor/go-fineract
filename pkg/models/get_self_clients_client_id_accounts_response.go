package models

// GetSelfClientsClientIdAccountsResponse GetSelfClientsClientIdAccountsResponse
type GetSelfClientsClientIdAccountsResponse struct {
	LoanAccounts    []GetSelfClientsLoanAccounts    `json:"loanAccounts,omitempty"`
	SavingsAccounts []GetSelfClientsSavingsAccounts `json:"savingsAccounts,omitempty"`
}
