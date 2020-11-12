package models

// PostLoansLoanIdChargesChargeIdResponse PostLoansLoanIdChargesChargeIdResponse
type PostLoansLoanIdChargesChargeIdResponse struct {
	OfficeId   int64 `json:"officeId,omitempty"`
	ClientId   int64 `json:"clientId,omitempty"`
	LoanId     int64 `json:"loanId,omitempty"`
	SavingsId  int64 `json:"savingsId,omitempty"`
	ResourceId int32 `json:"resourceId,omitempty"`
}
