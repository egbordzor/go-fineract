package models

// PutLoansLoanIdCollateralsCollateralIdResponse PutLoansLoanIdCollateralsCollateralIdResponse
type PutLoansLoanIdCollateralsCollateralIdResponse struct {
	LoanId     int32                                         `json:"loanId,omitempty"`
	ResourceId int32                                         `json:"resourceId,omitempty"`
	Changes    PutLoansLoandIdCollateralsCollateralIdRequest `json:"changes,omitempty"`
}
