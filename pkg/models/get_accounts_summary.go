package models

// GetAccountsSummary struct for GetAccountsSummary
type GetAccountsSummary struct {
	Id                            int32               `json:"id,omitempty"`
	AccountNo                     int64               `json:"accountNo,omitempty"`
	TotalApprovedShares           int32               `json:"totalApprovedShares,omitempty"`
	TotalPendingForApprovalShares int32               `json:"totalPendingForApprovalShares,omitempty"`
	ProductId                     int32               `json:"productId,omitempty"`
	ProductName                   string              `json:"productName,omitempty"`
	Status                        GetAccountsStatus   `json:"status,omitempty"`
	Timeline                      GetAccountsTimeline `json:"timeline,omitempty"`
	Currency                      GetAccountsCurrency `json:"currency,omitempty"`
}
