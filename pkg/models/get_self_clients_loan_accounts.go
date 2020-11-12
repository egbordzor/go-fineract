package models

// GetSelfClientsLoanAccounts struct for GetSelfClientsLoanAccounts
type GetSelfClientsLoanAccounts struct {
	Id          int32                            `json:"id,omitempty"`
	AccountNo   int64                            `json:"accountNo,omitempty"`
	ExternalId  int32                            `json:"externalId,omitempty"`
	ProductId   int32                            `json:"productId,omitempty"`
	ProductName string                           `json:"productName,omitempty"`
	Status      GetSelfClientsLoanAccountsStatus `json:"status,omitempty"`
	LoanType    GetSelfClientsLoanAccountsType   `json:"loanType,omitempty"`
	LoanCycle   int32                            `json:"loanCycle,omitempty"`
}
