package models

// PutSelfBeneficiariesChanges struct for PutSelfBeneficiariesChanges
type PutSelfBeneficiariesChanges struct {
	TransferLimit int32  `json:"transferLimit,omitempty"`
	Name          string `json:"name,omitempty"`
}
