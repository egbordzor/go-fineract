package models

// PostFinancialActivityAccountsRequest PostFinancialActivityAccountsRequest
type PostFinancialActivityAccountsRequest struct {
	FinancialActivityId int64 `json:"financialActivityId,omitempty"`
	GlAccountId         int64 `json:"glAccountId,omitempty"`
}
