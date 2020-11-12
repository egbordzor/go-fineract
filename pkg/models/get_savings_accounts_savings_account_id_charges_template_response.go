package models

// GetSavingsAccountsSavingsAccountIdChargesTemplateResponse GetSavingsAccountsSavingsAccountIdChargesTemplateResponse
type GetSavingsAccountsSavingsAccountIdChargesTemplateResponse struct {
	AmountPaid       float32                    `json:"amountPaid,omitempty"`
	AmountWaived     float32                    `json:"amountWaived,omitempty"`
	AmountWrittenOff float32                    `json:"amountWrittenOff,omitempty"`
	ChargeOptions    []GetSavingsChargesOptions `json:"chargeOptions,omitempty"`
	Penalty          bool                       `json:"penalty,omitempty"`
}
