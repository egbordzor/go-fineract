package models

// GetProductsTypeResponse GetProductsTypeResponse
type GetProductsTypeResponse struct {
	TotalFilteredRecords int32                  `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetProductsPageItems `json:"pageItems,omitempty"`
}
