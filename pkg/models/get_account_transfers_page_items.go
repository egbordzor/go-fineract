package models

// GetAccountTransfersPageItems struct for GetAccountTransfersPageItems
type GetAccountTransfersPageItems struct {
	Id                  int32                                     `json:"id,omitempty"`
	Reversed            bool                                      `json:"reversed,omitempty"`
	Currency            GetAccountTransfersPageItemsCurrency      `json:"currency,omitempty"`
	TransferAmount      float32                                   `json:"transferAmount,omitempty"`
	TransferDate        string                                    `json:"transferDate,omitempty"`
	TransferDescription string                                    `json:"transferDescription,omitempty"`
	FromOffice          GetAccountTransfersPageItemsFromOffice    `json:"fromOffice,omitempty"`
	FromClient          GetAccountTransfersFromClientOptions      `json:"fromClient,omitempty"`
	FromAccountType     GetAccountTransfersFromAccountType        `json:"fromAccountType,omitempty"`
	FromAccount         GetAccountTransfersPageItemsFromAccount   `json:"fromAccount,omitempty"`
	ToOffice            GetAccountTransfersPageItemsFromOffice    `json:"toOffice,omitempty"`
	ToClient            GetAccountTransfersFromClientOptions      `json:"toClient,omitempty"`
	ToAccountType       GetAccountTransfersPageItemsToAccountType `json:"toAccountType,omitempty"`
	ToAccount           GetAccountTransfersPageItemsFromAccount   `json:"toAccount,omitempty"`
}
