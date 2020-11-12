package models

// GetClientIdProductIdAccountingMappings struct for GetClientIdProductIdAccountingMappings
type GetClientIdProductIdAccountingMappings struct {
	ShareReferenceId       GetShareAccountsShareReferenceId       `json:"shareReferenceId,omitempty"`
	IncomeFromFeeAccountId GetShareAccountsIncomeFromFeeAccountId `json:"incomeFromFeeAccountId,omitempty"`
	ShareEquityId          GetShareAccountsShareEquityId          `json:"ShareEquityId,omitempty"`
	ShareSuspenseId        GetShareAccountsShareSuspenseId        `json:"shareSuspenseId,omitempty"`
}
