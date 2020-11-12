package models

// GetGroupsGroupIdAccountsSavingAccounts struct for GetGroupsGroupIdAccountsSavingAccounts
type GetGroupsGroupIdAccountsSavingAccounts struct {
	Id          int32                                     `json:"id,omitempty"`
	AccountNo   int64                                     `json:"accountNo,omitempty"`
	ProductId   int32                                     `json:"productId,omitempty"`
	ProductName string                                    `json:"productName,omitempty"`
	Status      GetGroupsGroupIdAccountsSavingStatus      `json:"status,omitempty"`
	Currency    GetGroupsGroupIdAccountsSavingCurrency    `json:"currency,omitempty"`
	AccountType GetGroupsGroupIdAccountsSavingAccountType `json:"accountType,omitempty"`
}
