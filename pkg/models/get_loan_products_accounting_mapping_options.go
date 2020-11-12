package models

// GetLoanProductsAccountingMappingOptions struct for GetLoanProductsAccountingMappingOptions
type GetLoanProductsAccountingMappingOptions struct {
	LiabilityAccountOptions []GetLoanProductsLiabilityAccountOptions `json:"liabilityAccountOptions,omitempty"`
	AssetAccountOptions     []GetLoanProductsAssetAccountOptions     `json:"assetAccountOptions,omitempty"`
	ExpenseAccountOptions   []GetLoanProductsExpenseAccountOptions   `json:"expenseAccountOptions,omitempty"`
	IncomeAccountOptions    []GetLoanProductsIncomeAccountOptions    `json:"incomeAccountOptions,omitempty"`
}
