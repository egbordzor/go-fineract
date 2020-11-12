package models

// GetAccountsPageItems struct for GetAccountsPageItems
type GetAccountsPageItems struct {
	Id              int32                            `json:"id,omitempty"`
	AccountNo       int64                            `json:"accountNo,omitempty"`
	ClientId        int32                            `json:"clientId,omitempty"`
	ClientName      string                           `json:"clientName,omitempty"`
	ProductId       int32                            `json:"productId,omitempty"`
	ProductName     string                           `json:"productName,omitempty"`
	Status          GetAccountsTypeStatus            `json:"status,omitempty"`
	Timeline        GetAccountsTypeTimeline          `json:"timeline,omitempty"`
	Currency        GetAccountsChargesCurrency       `json:"currency,omitempty"`
	PurchasedShares []GetAccountsTypePurchasedShares `json:"purchasedShares,omitempty"`
	Summary         GetAccountsTypeSummary           `json:"summary,omitempty"`
}
