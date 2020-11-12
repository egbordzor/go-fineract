package models

// PostRecurringDepositProductsChartSlabs struct for PostRecurringDepositProductsChartSlabs
type PostRecurringDepositProductsChartSlabs struct {
	Description        string  `json:"description,omitempty"`
	PeriodType         int32   `json:"periodType,omitempty"`
	FromPeriod         int32   `json:"fromPeriod,omitempty"`
	ToPeriod           int32   `json:"toPeriod,omitempty"`
	AnnualInterestRate float64 `json:"annualInterestRate,omitempty"`
}
