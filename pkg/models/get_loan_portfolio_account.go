package models

// GetLoanPortfolioAccount struct for GetLoanPortfolioAccount
type GetLoanPortfolioAccount struct {
	Id     int32  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	GlCode int32  `json:"glCode,omitempty"`
}
