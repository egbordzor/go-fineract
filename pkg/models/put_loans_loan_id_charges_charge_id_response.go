package models

// PutLoansLoanIdChargesChargeIdResponse PutLoansLoanIdChargesChargeIdResponse
type PutLoansLoanIdChargesChargeIdResponse struct {
	OfficeId   int64                                `json:"officeId,omitempty"`
	ClientId   int64                                `json:"clientId,omitempty"`
	LoanId     int64                                `json:"loanId,omitempty"`
	ResourceId int32                                `json:"resourceId,omitempty"`
	Changes    PutLoansLoanIdChargesChargeIdRequest `json:"changes,omitempty"`
}
