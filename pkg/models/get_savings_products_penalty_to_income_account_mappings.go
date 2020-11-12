package models

// GetSavingsProductsPenaltyToIncomeAccountMappings struct for GetSavingsProductsPenaltyToIncomeAccountMappings
type GetSavingsProductsPenaltyToIncomeAccountMappings struct {
	Charge        GetSavingsProductsPenaltyToIncomeAccountMappingsCharge `json:"charge,omitempty"`
	IncomeAccount GetSavingsProductsIncomeFromPenaltyAccount             `json:"incomeAccount,omitempty"`
}
