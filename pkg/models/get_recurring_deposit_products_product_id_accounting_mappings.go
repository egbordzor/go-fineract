package models

// GetRecurringDepositProductsProductIdAccountingMappings struct for GetRecurringDepositProductsProductIdAccountingMappings
type GetRecurringDepositProductsProductIdAccountingMappings struct {
	SavingsReferenceAccount    GetRecurringDepositProductsProductIdSavingsReferenceAccount    `json:"savingsReferenceAccount,omitempty"`
	IncomeFromFeeAccount       GetRecurringDepositProductsProductIdIncomeFromFeeAccount       `json:"incomeFromFeeAccount,omitempty"`
	IncomeFromPenaltyAccount   GetRecurringDepositProductsProductIdIncomeFromPenaltyAccount   `json:"incomeFromPenaltyAccount,omitempty"`
	InterestOnSavingsAccount   GetRecurringDepositProductsProductIdInterestOnSavingsAccount   `json:"interestOnSavingsAccount,omitempty"`
	SavingsControlAccount      GetRecurringDepositProductsProductIdSavingsControlAccount      `json:"savingsControlAccount,omitempty"`
	TransfersInSuspenseAccount GetRecurringDepositProductsProductIdTransfersInSuspenseAccount `json:"transfersInSuspenseAccount,omitempty"`
}
