package models

// GetClientsSavingsAccounts struct for GetClientsSavingsAccounts
type GetClientsSavingsAccounts struct {
	Id          int32                             `json:"id,omitempty"`
	AccountNo   int64                             `json:"accountNo,omitempty"`
	ProductId   int32                             `json:"productId,omitempty"`
	ProductName string                            `json:"productName,omitempty"`
	Status      GetClientsSavingsAccountsStatus   `json:"status,omitempty"`
	Currency    GetClientsSavingsAccountsCurrency `json:"currency,omitempty"`
}
