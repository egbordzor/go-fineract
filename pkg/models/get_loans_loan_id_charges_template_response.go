package models

// GetLoansLoanIdChargesTemplateResponse GetLoansLoanIdChargesTemplateResponse
type GetLoansLoanIdChargesTemplateResponse struct {
	AmountPaid       float32                              `json:"amountPaid,omitempty"`
	AmountWaived     float32                              `json:"amountWaived,omitempty"`
	AmountWrittenOff float32                              `json:"amountWrittenOff,omitempty"`
	ChargeOptions    []GetLoanChargeTemplateChargeOptions `json:"chargeOptions,omitempty"`
	Penalty          bool                                 `json:"penalty,omitempty"`
}
