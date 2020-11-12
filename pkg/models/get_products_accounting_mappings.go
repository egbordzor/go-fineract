package models

// GetProductsAccountingMappings struct for GetProductsAccountingMappings
type GetProductsAccountingMappings struct {
	ShareReferenceId       GetShareReferenceId       `json:"shareReferenceId,omitempty"`
	IncomeFromFeeAccountId GetIncomeFromFeeAccountId `json:"incomeFromFeeAccountId,omitempty"`
	ShareEquityId          GetShareEquityId          `json:"shareEquityId,omitempty"`
	ShareSuspenseId        GetShareSuspenseId        `json:"shareSuspenseId,omitempty"`
}
