package models

// GetChargesTemplateResponse GetChargesTemplateResponse
type GetChargesTemplateResponse struct {
	Active                              string                                               `json:"active,omitempty"`
	Penalty                             string                                               `json:"penalty,omitempty"`
	CurrencyOptions                     []GetChargesCurrencyResponse                         `json:"currencyOptions,omitempty"`
	ChargeCalculationTypeOptions        []GetChargesCalculationTypeResponse                  `json:"chargeCalculationTypeOptions,omitempty"`
	ChargeAppliesToOptions              []GetChargesAppliesToResponse                        `json:"chargeAppliesToOptions,omitempty"`
	ChargeTimeTypeOptions               []GetChargesTimeTypeResponse                         `json:"chargeTimeTypeOptions,omitempty"`
	ChargePaymentModeOptions            []GetChargesPaymentModeResponse                      `json:"chargePaymentModeOptions,omitempty"`
	LoanChargeCalculationTypeOptions    []GetChargesTemplateLoanChargeCalculationTypeOptions `json:"loanChargeCalculationTypeOptions,omitempty"`
	LoanChargeTimeTypeOptions           []GetChargesTemplateLoanChargeTimeTypeOptions        `json:"loanChargeTimeTypeOptions,omitempty"`
	SavingsChargeCalculationTypeOptions []GetChargesTemplateLoanChargeCalculationTypeOptions `json:"savingsChargeCalculationTypeOptions,omitempty"`
	SavingsChargeTimeTypeOptions        []GetChargesTemplateLoanChargeTimeTypeOptions        `json:"savingsChargeTimeTypeOptions,omitempty"`
	FeeFrequencyOptions                 []GetChargesTemplateFeeFrequencyOptions              `json:"feeFrequencyOptions,omitempty"`
}
