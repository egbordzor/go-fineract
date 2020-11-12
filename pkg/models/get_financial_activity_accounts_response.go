package models

// GetFinancialActivityAccountsResponse GetFinancialActivityAccountsResponse
type GetFinancialActivityAccountsResponse struct {
	Id                    int64                 `json:"id,omitempty"`
	FinancialActivityData FinancialActivityData `json:"financialActivityData,omitempty"`
	GlAccountData         GlAccountData         `json:"glAccountData,omitempty"`
}
