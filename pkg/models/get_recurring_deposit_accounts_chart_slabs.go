package models

// GetRecurringDepositAccountsChartSlabs struct for GetRecurringDepositAccountsChartSlabs
type GetRecurringDepositAccountsChartSlabs struct {
	Id                 int32                                           `json:"id,omitempty"`
	PeriodType         GetRecurringDepositAccountsPeriodType           `json:"periodType,omitempty"`
	FromPeriod         int32                                           `json:"fromPeriod,omitempty"`
	ToPeriod           int32                                           `json:"toPeriod,omitempty"`
	AnnualInterestRate float64                                         `json:"annualInterestRate,omitempty"`
	Currency           GetRecurringDepositAccountsAccountChartCurrency `json:"currency,omitempty"`
}
