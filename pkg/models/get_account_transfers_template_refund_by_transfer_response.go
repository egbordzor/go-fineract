package models

// GetAccountTransfersTemplateRefundByTransferResponse GetAccountTransfersTemplateRefundByTransferResponse
type GetAccountTransfersTemplateRefundByTransferResponse struct {
	Currency               GetAccountTransfersTemplateRefundByTransferCurrency             `json:"currency,omitempty"`
	TransferAmount         float32                                                         `json:"transferAmount,omitempty"`
	TransferDate           string                                                          `json:"transferDate,omitempty"`
	FromOffice             GetAccountTransfersTemplateRefundByTransferFromOffice           `json:"fromOffice,omitempty"`
	FromClient             GetAccountTransfersTemplateRefundByTransferFromClient           `json:"fromClient,omitempty"`
	FromAccountType        GetAccountTransfersPageItemsToAccountType                       `json:"fromAccountType,omitempty"`
	FromAccount            GetAccountTransfersTemplateRefundByTransferFromAccount          `json:"fromAccount,omitempty"`
	ToOffice               GetAccountTransfersTemplateRefundByTransferFromOffice           `json:"toOffice,omitempty"`
	ToClient               GetAccountTransfersTemplateRefundByTransferToClient             `json:"toClient,omitempty"`
	ToAccountType          GetAccountTransfersFromAccountType                              `json:"toAccountType,omitempty"`
	ToAccount              GetAccountTransfersTemplateRefundByTransferToAccount            `json:"toAccount,omitempty"`
	FromOfficeOptions      []GetAccountTransfersTemplateRefundByTransferFromOfficeOptions  `json:"fromOfficeOptions,omitempty"`
	FromClientOptions      []GetAccountTransfersTemplateRefundByTransferFromClientOptions  `json:"fromClientOptions,omitempty"`
	FromAccountTypeOptions []GetAccountTransfersFromAccountType                            `json:"fromAccountTypeOptions,omitempty"`
	FromAccountOptions     []GetAccountTransfersTemplateRefundByTransferFromAccountOptions `json:"fromAccountOptions,omitempty"`
	ToOfficeOptions        []GetAccountTransfersTemplateRefundByTransferFromOfficeOptions  `json:"toOfficeOptions,omitempty"`
	ToClientOptions        []GetAccountTransfersTemplateRefundByTransferFromClientOptions  `json:"toClientOptions,omitempty"`
	ToAccountTypeOptions   []GetAccountTransfersFromAccountType                            `json:"toAccountTypeOptions,omitempty"`
	ToAccountOptions       []GetAccountTransfersTemplateRefundByTransferToAccount          `json:"toAccountOptions,omitempty"`
}
