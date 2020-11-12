package models

// GetProductsMarketPrice struct for GetProductsMarketPrice
type GetProductsMarketPrice struct {
	Id         int32  `json:"id,omitempty"`
	FromDate   string `json:"fromDate,omitempty"`
	ShareValue int32  `json:"shareValue,omitempty"`
}
