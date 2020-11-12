package models

// GetLoanFundSourceAccount struct for GetLoanFundSourceAccount
type GetLoanFundSourceAccount struct {
	Id     int32  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	GlCode int32  `json:"glCode,omitempty"`
}
