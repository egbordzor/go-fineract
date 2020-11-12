package models

// GetLoansLoanIdCollateralsResponse GetLoansLoanIdCollateralsResponse
type GetLoansLoanIdCollateralsResponse struct {
	Id          int32                         `json:"id,omitempty"`
	Type        GetCollateralTypeResponse     `json:"type,omitempty"`
	Value       int64                         `json:"value,omitempty"`
	Description string                        `json:"description,omitempty"`
	Currency    GetCollateralCurrencyResponse `json:"currency,omitempty"`
}
