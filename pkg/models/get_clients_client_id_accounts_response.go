package models

// GetClientsClientIdAccountsResponse GetClientsClientIdAccountsResponse
type GetClientsClientIdAccountsResponse struct {
	LoanAccounts    []GetClientsLoanAccounts    `json:"loanAccounts,omitempty"`
	SavingsAccounts []GetClientsSavingsAccounts `json:"savingsAccounts,omitempty"`
}
