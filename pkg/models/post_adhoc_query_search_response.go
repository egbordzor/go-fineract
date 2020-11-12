package models

// PostAdhocQuerySearchResponse PostAdhocQuerySearchResponse
type PostAdhocQuerySearchResponse struct {
	OfficeName      string `json:"officeName,omitempty"`
	LoanProductName string `json:"loanProductName,omitempty"`
	LoanOutStanding int64  `json:"loanOutStanding,omitempty"`
	Percentage      int64  `json:"percentage,omitempty"`
}
