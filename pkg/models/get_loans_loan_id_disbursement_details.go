package models

// GetLoansLoanIdDisbursementDetails struct for GetLoansLoanIdDisbursementDetails
type GetLoansLoanIdDisbursementDetails struct {
	Id                       int32   `json:"id,omitempty"`
	ExpectedDisbursementDate string  `json:"expectedDisbursementDate,omitempty"`
	Principal                float32 `json:"principal,omitempty"`
	ApprovedPrincipal        float32 `json:"approvedPrincipal,omitempty"`
}
