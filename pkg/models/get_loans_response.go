package models

// GetLoansResponse GetLoansResponse
type GetLoansResponse struct {
	TotalFilteredRecords int32                    `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetLoansLoanIdResponse `json:"pageItems,omitempty"`
}
