package models

// GetFixedDepositAccountsAccountIdResponse GetFixedDepositAccountsAccountIdResponse
type GetFixedDepositAccountsAccountIdResponse struct {
	Id                                int32                                                    `json:"id,omitempty"`
	AccountNo                         int64                                                    `json:"accountNo,omitempty"`
	ExternalId                        string                                                   `json:"externalId,omitempty"`
	ClientId                          int32                                                    `json:"clientId,omitempty"`
	ClientName                        string                                                   `json:"clientName,omitempty"`
	SavingsProductId                  int32                                                    `json:"savingsProductId,omitempty"`
	SavingsProductName                string                                                   `json:"savingsProductName,omitempty"`
	FieldOfficerId                    int32                                                    `json:"fieldOfficerId,omitempty"`
	Status                            GetFixedDepositAccountsStatus                            `json:"status,omitempty"`
	Timeline                          GetFixedDepositAccountsTimeline                          `json:"timeline,omitempty"`
	Currency                          GetFixedDepositAccountsAccountIdCurrency                 `json:"currency,omitempty"`
	InterestCompoundingPeriodType     GetFixedDepositAccountsInterestCompoundingPeriodType     `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         GetFixedDepositAccountsInterestPostingPeriodType         `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           GetFixedDepositAccountsInterestCalculationType           `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType GetFixedDepositAccountsInterestCalculationDaysInYearType `json:"interestCalculationDaysInYearType,omitempty"`
	InterestFreePeriodApplicable      bool                                                     `json:"interestFreePeriodApplicable,omitempty"`
	PreClosurePenalApplicable         bool                                                     `json:"preClosurePenalApplicable,omitempty"`
	MinDepositTerm                    int32                                                    `json:"minDepositTerm,omitempty"`
	MaxDepositTerm                    int32                                                    `json:"maxDepositTerm,omitempty"`
	MinDepositTermType                GetFixedDepositAccountsMinDepositTermType                `json:"minDepositTermType,omitempty"`
	MaxDepositTermType                GetFixedDepositAccountsMaxDepositTermType                `json:"maxDepositTermType,omitempty"`
	DepositAmount                     float32                                                  `json:"depositAmount,omitempty"`
	MaturityAmount                    float32                                                  `json:"maturityAmount,omitempty"`
	MaturityDate                      string                                                   `json:"maturityDate,omitempty"`
	DepositPeriod                     int32                                                    `json:"depositPeriod,omitempty"`
	DepositPeriodFrequency            GetFixedDepositAccountsDepositPeriodFrequency            `json:"depositPeriodFrequency,omitempty"`
	Summary                           GetFixedDepositAccountsAccountIdSummary                  `json:"summary,omitempty"`
	AccountChart                      GetFixedDepositAccountsAccountChart                      `json:"accountChart,omitempty"`
}
