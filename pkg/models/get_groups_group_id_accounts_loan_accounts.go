package models

// GetGroupsGroupIdAccountsLoanAccounts struct for GetGroupsGroupIdAccountsLoanAccounts
type GetGroupsGroupIdAccountsLoanAccounts struct {
	Id          int32                            `json:"id,omitempty"`
	AccountNo   int64                            `json:"accountNo,omitempty"`
	ProductId   int32                            `json:"productId,omitempty"`
	ProductName string                           `json:"productName,omitempty"`
	Status      GetGroupsGroupIdAccountsStatus   `json:"status,omitempty"`
	LoanType    GetGroupsGroupIdAccountsLoanType `json:"loanType,omitempty"`
}
