package models

// GetGroupsGroupIdAccountsMemberLoanAccounts struct for GetGroupsGroupIdAccountsMemberLoanAccounts
type GetGroupsGroupIdAccountsMemberLoanAccounts struct {
	Id          int32                                    `json:"id,omitempty"`
	AccountNo   int64                                    `json:"accountNo,omitempty"`
	ProductId   int32                                    `json:"productId,omitempty"`
	ProductName string                                   `json:"productName,omitempty"`
	Status      GetGroupsGroupIdAccountsMemberLoanStatus `json:"status,omitempty"`
	LoanType    GetGroupsGroupIdAccountsMemberLoanType   `json:"loanType,omitempty"`
}
