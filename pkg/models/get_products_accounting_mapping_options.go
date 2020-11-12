package models

// GetProductsAccountingMappingOptions struct for GetProductsAccountingMappingOptions
type GetProductsAccountingMappingOptions struct {
	LiabilityAccountOptions []GetProductsLiabilityAccountOptions `json:"liabilityAccountOptions,omitempty"`
	AssetAccountOptions     []GetProductsAssetAccountOptions     `json:"assetAccountOptions,omitempty"`
	IncomeAccountOptions    []GetProductsIncomeAccountOptions    `json:"incomeAccountOptions,omitempty"`
	EquityAccountOptions    []GetProductsEquityAccountOptions    `json:"equityAccountOptions,omitempty"`
}
