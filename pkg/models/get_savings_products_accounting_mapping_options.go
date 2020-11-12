package models

// GetSavingsProductsAccountingMappingOptions struct for GetSavingsProductsAccountingMappingOptions
type GetSavingsProductsAccountingMappingOptions struct {
	LiabilityAccountOptions []GetSavingsProductsLiabilityAccountOptions `json:"liabilityAccountOptions,omitempty"`
	AssetAccountOptions     []GetSavingsProductsAssetAccountOptions     `json:"assetAccountOptions,omitempty"`
	ExpenseAccountOptions   []GetSavingsProductsExpenseAccountOptions   `json:"expenseAccountOptions,omitempty"`
	IncomeAccountOptions    []GetSavingsProductsIncomeAccountOptions    `json:"incomeAccountOptions,omitempty"`
}
