package models

// PutStandingInstructionsStandingInstructionIdResponse PutStandingInstructionsStandingInstructionIdResponse
type PutStandingInstructionsStandingInstructionIdResponse struct {
	ResourceId int32                               `json:"resourceId,omitempty"`
	Changes    PutUpdateStandingInstructionChanges `json:"changes,omitempty"`
}
