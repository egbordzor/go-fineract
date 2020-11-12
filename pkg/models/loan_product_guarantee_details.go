package models

// LoanProductGuaranteeDetails struct for LoanProductGuaranteeDetails
type LoanProductGuaranteeDetails struct {
	Id                            int64   `json:"id,omitempty"`
	MandatoryGuarantee            float32 `json:"mandatoryGuarantee,omitempty"`
	MinimumGuaranteeFromOwnFunds  float32 `json:"minimumGuaranteeFromOwnFunds,omitempty"`
	MinimumGuaranteeFromGuarantor float32 `json:"minimumGuaranteeFromGuarantor,omitempty"`
	New                           bool    `json:"new,omitempty"`
}
