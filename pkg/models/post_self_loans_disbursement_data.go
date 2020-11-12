package models

// PostSelfLoansDisbursementData struct for PostSelfLoansDisbursementData
type PostSelfLoansDisbursementData struct {
	ExpectedDisbursementDate string `json:"expectedDisbursementDate,omitempty"`
	Principal                int64  `json:"principal,omitempty"`
	ApprovedPrincipal        int64  `json:"approvedPrincipal,omitempty"`
}
