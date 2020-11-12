package models

// GetLoanAccountingMappings struct for GetLoanAccountingMappings
type GetLoanAccountingMappings struct {
	FundSourceAccount           GetLoanFundSourceAccount           `json:"fundSourceAccount,omitempty"`
	LoanPortfolioAccount        GetLoanPortfolioAccount            `json:"loanPortfolioAccount,omitempty"`
	TransfersInSuspenseAccount  GetLoanTransfersInSuspenseAccount  `json:"transfersInSuspenseAccount,omitempty"`
	InterestOnLoanAccount       GetLoanInterestOnLoanAccount       `json:"interestOnLoanAccount,omitempty"`
	IncomeFromFeeAccount        GetLoanIncomeFromFeeAccount        `json:"incomeFromFeeAccount,omitempty"`
	IncomeFromPenaltyAccount    GetLoanIncomeFromPenaltyAccount    `json:"incomeFromPenaltyAccount,omitempty"`
	WriteOffAccount             GetLoanWriteOffAccount             `json:"writeOffAccount,omitempty"`
	OverpaymentLiabilityAccount GetLoanOverpaymentLiabilityAccount `json:"overpaymentLiabilityAccount,omitempty"`
}
