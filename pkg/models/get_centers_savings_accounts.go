package models

// GetCentersSavingsAccounts struct for GetCentersSavingsAccounts
type GetCentersSavingsAccounts struct {
	Id          int32                      `json:"id,omitempty"`
	AccountNo   int64                      `json:"accountNo,omitempty"`
	ProductId   int32                      `json:"productId,omitempty"`
	ProductName string                     `json:"productName,omitempty"`
	Status      GetCentersCenterIdStatus   `json:"status,omitempty"`
	Currency    GetCentersCenterIdCurrency `json:"currency,omitempty"`
	AccountType GetCentersAccountType      `json:"accountType,omitempty"`
	Timeline    GetCentersTimeline         `json:"timeline,omitempty"`
	DepositType GetCentersDepositType      `json:"depositType,omitempty"`
}
