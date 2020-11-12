package models

// GetGroupsGroupIdAccountsResponse GetGroupsGroupIdAccountsResponse
type GetGroupsGroupIdAccountsResponse struct {
	LoanAccounts          []GetGroupsGroupIdAccountsLoanAccounts          `json:"loanAccounts,omitempty"`
	SavingsAccounts       []GetGroupsGroupIdAccountsSavingAccounts        `json:"savingsAccounts,omitempty"`
	MemberLoanAccounts    []GetGroupsGroupIdAccountsMemberLoanAccounts    `json:"memberLoanAccounts,omitempty"`
	MemberSavingsAccounts []GetGroupsGroupIdAccountsMemberSavingsAccounts `json:"memberSavingsAccounts,omitempty"`
}
