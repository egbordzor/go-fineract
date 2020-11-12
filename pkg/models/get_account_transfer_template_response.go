package models

// GetAccountTransferTemplateResponse GetAccountTransferTemplateResponse
type GetAccountTransferTemplateResponse struct {
	AccountTypeOptions     []GetAccountOptions     `json:"accountTypeOptions,omitempty"`
	FromAccountTypeOptions []GetFromAccountOptions `json:"fromAccountTypeOptions,omitempty"`
	ToAccountTypeOptions   []GetFromAccountOptions `json:"toAccountTypeOptions,omitempty"`
}
