package models

// PutSelfBeneficiariesTptBeneficiaryIdResponse PutSelfBeneficiariesTPTBeneficiaryIdResponse
type PutSelfBeneficiariesTptBeneficiaryIdResponse struct {
	ResourceId int32                       `json:"resourceId,omitempty"`
	Changes    PutSelfBeneficiariesChanges `json:"changes,omitempty"`
}
