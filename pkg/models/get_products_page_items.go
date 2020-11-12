package models

// GetProductsPageItems struct for GetProductsPageItems
type GetProductsPageItems struct {
	Id          int32  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ShortName   string `json:"shortName,omitempty"`
	TotalShares int32  `json:"totalShares,omitempty"`
}
