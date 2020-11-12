package models

// GetAccountsTypePurchasedShares struct for GetAccountsTypePurchasedShares
type GetAccountsTypePurchasedShares struct {
	Id             int32  `json:"id,omitempty"`
	PurchasedDate  string `json:"purchasedDate,omitempty"`
	NumberOfShares int32  `json:"numberOfShares,omitempty"`
	PurchasedPrice int32  `json:"purchasedPrice,omitempty"`
}
