package models

// GetClientsLoanAccounts struct for GetClientsLoanAccounts
type GetClientsLoanAccounts struct {
	Id          int32                        `json:"id,omitempty"`
	AccountNo   int64                        `json:"accountNo,omitempty"`
	ExternalId  int32                        `json:"externalId,omitempty"`
	ProductId   int32                        `json:"productId,omitempty"`
	ProductName string                       `json:"productName,omitempty"`
	Status      GetClientsLoanAccountsStatus `json:"status,omitempty"`
	LoanType    GetClientsLoanAccountsType   `json:"loanType,omitempty"`
	LoanCycle   int32                        `json:"loanCycle,omitempty"`
}
