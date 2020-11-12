package models

// GetSelfSavingsAccountsResponse GetSelfSavingsAccountsResponse
type GetSelfSavingsAccountsResponse struct {
	Id                                int32                                           `json:"id,omitempty"`
	AccountNo                         int64                                           `json:"accountNo,omitempty"`
	ClientId                          int32                                           `json:"clientId,omitempty"`
	ClientName                        string                                          `json:"clientName,omitempty"`
	SavingsProductId                  int32                                           `json:"savingsProductId,omitempty"`
	SavingsProductName                string                                          `json:"savingsProductName,omitempty"`
	FieldOfficerId                    int32                                           `json:"fieldOfficerId,omitempty"`
	Status                            GetSelfSavingsStatus                            `json:"status,omitempty"`
	Timeline                          GetSelfSavingsTimeline                          `json:"timeline,omitempty"`
	Currency                          GetSelfSavingsCurrency                          `json:"currency,omitempty"`
	NominalAnnualInterestRate         float64                                         `json:"nominalAnnualInterestRate,omitempty"`
	InterestCompoundingPeriodType     GetSelfSavingsInterestCompoundingPeriodType     `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         GetSelfSavingsInterestPostingPeriodType         `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           GetSelfSavingsInterestCalculationType           `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType GetSelfSavingsInterestCalculationDaysInYearType `json:"interestCalculationDaysInYearType,omitempty"`
	Summary                           GetSelfSavingsSummary                           `json:"summary,omitempty"`
}
