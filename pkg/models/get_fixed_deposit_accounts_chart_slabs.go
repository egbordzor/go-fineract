package models

// GetFixedDepositAccountsChartSlabs struct for GetFixedDepositAccountsChartSlabs
type GetFixedDepositAccountsChartSlabs struct {
	Id                 int32                                       `json:"id,omitempty"`
	PeriodType         GetFixedDepositAccountsPeriodType           `json:"periodType,omitempty"`
	FromPeriod         int32                                       `json:"fromPeriod,omitempty"`
	ToPeriod           int32                                       `json:"toPeriod,omitempty"`
	AnnualInterestRate float64                                     `json:"annualInterestRate,omitempty"`
	Currency           GetFixedDepositAccountsAccountChartCurrency `json:"currency,omitempty"`
}
