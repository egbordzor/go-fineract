package models

// PutSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse PutSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse
type PutSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdResponse struct {
	OfficeId   int32             `json:"officeId,omitempty"`
	ClientId   int32             `json:"clientId,omitempty"`
	SavingsId  int32             `json:"savingsId,omitempty"`
	ResourceId int32             `json:"resourceId,omitempty"`
	Changes    PutSavingsChanges `json:"changes,omitempty"`
}
