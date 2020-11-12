package models

// GetSavingsPageItems struct for GetSavingsPageItems
type GetSavingsPageItems struct {
	Id                                int32                                       `json:"id,omitempty"`
	AccountNo                         string                                      `json:"accountNo,omitempty"`
	ClientId                          int32                                       `json:"clientId,omitempty"`
	ClientName                        string                                      `json:"clientName,omitempty"`
	SavingsProductId                  int32                                       `json:"savingsProductId,omitempty"`
	SavingsProductName                string                                      `json:"savingsProductName,omitempty"`
	FieldOfficerId                    int32                                       `json:"fieldOfficerId,omitempty"`
	Status                            GetSavingsStatus                            `json:"status,omitempty"`
	Timeline                          GetSavingsTimeline                          `json:"timeline,omitempty"`
	Currency                          GetSavingsCurrency                          `json:"currency,omitempty"`
	NominalAnnualInterestRate         float64                                     `json:"nominalAnnualInterestRate,omitempty"`
	InterestCompoundingPeriodType     GetSavingsInterestCompoundingPeriodType     `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         GetSavingsInterestPostingPeriodType         `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           GetSavingsInterestCalculationType           `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType GetSavingsInterestCalculationDaysInYearType `json:"interestCalculationDaysInYearType,omitempty"`
	Summary                           GetSavingsSummary                           `json:"summary,omitempty"`
}
