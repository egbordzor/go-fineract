package models

// GetFixedDepositProductsProductIdPenaltyToIncomeAccountMappings struct for GetFixedDepositProductsProductIdPenaltyToIncomeAccountMappings
type GetFixedDepositProductsProductIdPenaltyToIncomeAccountMappings struct {
	Charge        GetFixedDepositProductsProductIdPenaltyToIncomeAccountMappingsCharge `json:"charge,omitempty"`
	IncomeAccount GetFixedDepositProductsProductIdIncomeFromPenaltyAccount             `json:"incomeAccount,omitempty"`
}
