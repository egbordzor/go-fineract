package models

// GetSavingsProductsAccountingMappings struct for GetSavingsProductsAccountingMappings
type GetSavingsProductsAccountingMappings struct {
	SavingsReferenceAccount    GetSavingsProductsSavingsReferenceAccount    `json:"savingsReferenceAccount,omitempty"`
	IncomeFromFeeAccount       GetSavingsProductsIncomeFromFeeAccount       `json:"incomeFromFeeAccount,omitempty"`
	IncomeFromPenaltyAccount   GetSavingsProductsIncomeFromPenaltyAccount   `json:"incomeFromPenaltyAccount,omitempty"`
	InterestOnSavingsAccount   GetSavingsProductsInterestOnSavingsAccount   `json:"interestOnSavingsAccount,omitempty"`
	SavingsControlAccount      GetSavingsProductsSavingsControlAccount      `json:"savingsControlAccount,omitempty"`
	TransfersInSuspenseAccount GetSavingsProductsTransfersInSuspenseAccount `json:"transfersInSuspenseAccount,omitempty"`
}
