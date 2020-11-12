package models

// GetRecurringDepositAccountsAccountIdResponse GetRecurringDepositAccountsAccountIdResponse
type GetRecurringDepositAccountsAccountIdResponse struct {
	Id                                int32                                                        `json:"id,omitempty"`
	AccountNo                         int64                                                        `json:"accountNo,omitempty"`
	ExternalId                        string                                                       `json:"externalId,omitempty"`
	ClientId                          int32                                                        `json:"clientId,omitempty"`
	ClientName                        string                                                       `json:"clientName,omitempty"`
	SavingsProductId                  int32                                                        `json:"savingsProductId,omitempty"`
	SavingsProductName                string                                                       `json:"savingsProductName,omitempty"`
	FieldOfficerId                    int32                                                        `json:"fieldOfficerId,omitempty"`
	Status                            GetRecurringDepositAccountsStatus                            `json:"status,omitempty"`
	Timeline                          GetRecurringDepositAccountsTimeline                          `json:"timeline,omitempty"`
	Currency                          GetRecurringDepositAccountsCurrency                          `json:"currency,omitempty"`
	InterestCompoundingPeriodType     GetRecurringDepositAccountsInterestCompoundingPeriodType     `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         GetRecurringDepositAccountsInterestPostingPeriodType         `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           GetRecurringDepositAccountsInterestCalculationType           `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType GetRecurringDepositAccountsInterestCalculationDaysInYearType `json:"interestCalculationDaysInYearType,omitempty"`
	PreClosurePenalApplicable         bool                                                         `json:"preClosurePenalApplicable,omitempty"`
	MinDepositTerm                    int32                                                        `json:"minDepositTerm,omitempty"`
	MaxDepositTerm                    int32                                                        `json:"maxDepositTerm,omitempty"`
	MinDepositTermType                GetRecurringDepositAccountsMinDepositTermType                `json:"minDepositTermType,omitempty"`
	MaxDepositTermType                GetRecurringDepositAccountsMaxDepositTermType                `json:"maxDepositTermType,omitempty"`
	RecurringDepositAmount            int32                                                        `json:"recurringDepositAmount,omitempty"`
	RecurringDepositFrequency         int32                                                        `json:"recurringDepositFrequency,omitempty"`
	ExpectedFirstDepositOnDate        string                                                       `json:"expectedFirstDepositOnDate,omitempty"`
	RecurringDepositFrequencyType     GetRecurringDepositAccountsRecurringDepositFrequencyType     `json:"recurringDepositFrequencyType,omitempty"`
	DepositPeriod                     int32                                                        `json:"depositPeriod,omitempty"`
	DepositPeriodFrequency            GetRecurringDepositAccountsDepositPeriodFrequency            `json:"depositPeriodFrequency,omitempty"`
	Summary                           GetRecurringDepositAccountsSummary                           `json:"summary,omitempty"`
	AccountChart                      GetRecurringDepositAccountsAccountChart                      `json:"accountChart,omitempty"`
}
