package models

// GetLoanFeeToIncomeAccountMappings struct for GetLoanFeeToIncomeAccountMappings
type GetLoanFeeToIncomeAccountMappings struct {
	Charge        GetLoanCharge               `json:"charge,omitempty"`
	IncomeAccount GetLoanIncomeFromFeeAccount `json:"incomeAccount,omitempty"`
}
