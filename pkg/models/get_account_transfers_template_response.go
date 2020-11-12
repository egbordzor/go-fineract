package models

// GetAccountTransfersTemplateResponse GetAccountTransfersTemplateResponse
type GetAccountTransfersTemplateResponse struct {
	TransferAmount         int64                                       `json:"transferAmount,omitempty"`
	TransferDate           string                                      `json:"transferDate,omitempty"`
	FromOffice             GetAccountTransfersFromOffice               `json:"fromOffice,omitempty"`
	FromAccountType        GetAccountTransfersFromAccountType          `json:"fromAccountType,omitempty"`
	FromOfficeOptions      []GetAccountTransfersFromOfficeOptions      `json:"fromOfficeOptions,omitempty"`
	FromClientOptions      []GetAccountTransfersFromClientOptions      `json:"fromClientOptions,omitempty"`
	FromAccountTypeOptions []GetAccountTransfersFromAccountTypeOptions `json:"fromAccountTypeOptions,omitempty"`
	ToOfficeOptions        []GetAccountTransfersToOfficeOptions        `json:"toOfficeOptions,omitempty"`
	ToAccountTypeOptions   []GetAccountTransfersToAccountTypeOptions   `json:"toAccountTypeOptions,omitempty"`
}
