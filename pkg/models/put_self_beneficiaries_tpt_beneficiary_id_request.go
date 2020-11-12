package models

// PutSelfBeneficiariesTptBeneficiaryIdRequest PutSelfBeneficiariesTPTBeneficiaryIdRequest
type PutSelfBeneficiariesTptBeneficiaryIdRequest struct {
	Name          string `json:"name,omitempty"`
	TransferLimit int32  `json:"transferLimit,omitempty"`
}
