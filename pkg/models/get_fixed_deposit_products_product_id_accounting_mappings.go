package models

// GetFixedDepositProductsProductIdAccountingMappings struct for GetFixedDepositProductsProductIdAccountingMappings
type GetFixedDepositProductsProductIdAccountingMappings struct {
	SavingsReferenceAccount    GetFixedDepositProductsProductIdSavingsReferenceAccount    `json:"savingsReferenceAccount,omitempty"`
	IncomeFromFeeAccount       GetFixedDepositProductsProductIdIncomeFromFeeAccount       `json:"incomeFromFeeAccount,omitempty"`
	IncomeFromPenaltyAccount   GetFixedDepositProductsProductIdIncomeFromPenaltyAccount   `json:"incomeFromPenaltyAccount,omitempty"`
	InterestOnSavingsAccount   GetFixedDepositProductsProductIdInterestOnSavingsAccount   `json:"interestOnSavingsAccount,omitempty"`
	SavingsControlAccount      GetFixedDepositProductsProductIdSavingsControlAccount      `json:"savingsControlAccount,omitempty"`
	TransfersInSuspenseAccount GetFixedDepositProductsProductIdTransfersInSuspenseAccount `json:"transfersInSuspenseAccount,omitempty"`
}
