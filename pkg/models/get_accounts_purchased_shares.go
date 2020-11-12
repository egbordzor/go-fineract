package models

// GetAccountsPurchasedShares struct for GetAccountsPurchasedShares
type GetAccountsPurchasedShares struct {
	Id             int32                            `json:"id,omitempty"`
	AccountId      int32                            `json:"accountId,omitempty"`
	PurchasedDate  string                           `json:"purchasedDate,omitempty"`
	NumberOfShares int32                            `json:"numberOfShares,omitempty"`
	PurchasedPrice float64                          `json:"purchasedPrice,omitempty"`
	Status         GetAccountsPurchasedSharesStatus `json:"status,omitempty"`
	Type           GetAccountsPurchasedSharesType   `json:"type,omitempty"`
	Amount         float64                          `json:"amount,omitempty"`
	ChargeAmount   float64                          `json:"chargeAmount,omitempty"`
	AmountPaid     float64                          `json:"amountPaid,omitempty"`
}
