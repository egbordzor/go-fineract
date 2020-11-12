package models

// GetAccountsTypeProductOptions struct for GetAccountsTypeProductOptions
type GetAccountsTypeProductOptions struct {
	Id          int32  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ShortName   string `json:"shortName,omitempty"`
	TotalShares int64  `json:"totalShares,omitempty"`
}
