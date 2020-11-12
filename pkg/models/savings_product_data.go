package models

// SavingsProductData struct for SavingsProductData
type SavingsProductData struct {
	Id                                int64          `json:"id,omitempty"`
	Name                              string         `json:"name,omitempty"`
	Currency                          CurrencyData   `json:"currency,omitempty"`
	NominalAnnualInterestRate         float32        `json:"nominalAnnualInterestRate,omitempty"`
	InterestCompoundingPeriodType     EnumOptionData `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         EnumOptionData `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           EnumOptionData `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType EnumOptionData `json:"interestCalculationDaysInYearType,omitempty"`
	MinRequiredOpeningBalance         float32        `json:"minRequiredOpeningBalance,omitempty"`
	LockinPeriodFrequency             int32          `json:"lockinPeriodFrequency,omitempty"`
	LockinPeriodFrequencyType         EnumOptionData `json:"lockinPeriodFrequencyType,omitempty"`
	WithdrawalFeeForTransfers         bool           `json:"withdrawalFeeForTransfers,omitempty"`
	AllowOverdraft                    bool           `json:"allowOverdraft,omitempty"`
	OverdraftLimit                    float32        `json:"overdraftLimit,omitempty"`
	MinRequiredBalance                float32        `json:"minRequiredBalance,omitempty"`
	DepositAccountType                string         `json:"depositAccountType,omitempty"`
}
