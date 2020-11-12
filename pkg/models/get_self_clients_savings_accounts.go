package models

// GetSelfClientsSavingsAccounts struct for GetSelfClientsSavingsAccounts
type GetSelfClientsSavingsAccounts struct {
	Id          int32                                 `json:"id,omitempty"`
	AccountNo   int64                                 `json:"accountNo,omitempty"`
	ProductId   int32                                 `json:"productId,omitempty"`
	ProductName string                                `json:"productName,omitempty"`
	Status      GetSelfClientsSavingsAccountsStatus   `json:"status,omitempty"`
	Currency    GetSelfClientsSavingsAccountsCurrency `json:"currency,omitempty"`
}
