package models

// GetRecurringDepositProductsProductIdChartSlabs struct for GetRecurringDepositProductsProductIdChartSlabs
type GetRecurringDepositProductsProductIdChartSlabs struct {
	Id                 int32                                          `json:"id,omitempty"`
	Description        string                                         `json:"description,omitempty"`
	PeriodType         GetRecurringDepositProductsProductIdPeriodType `json:"periodType,omitempty"`
	FromPeriod         int32                                          `json:"fromPeriod,omitempty"`
	ToPeriod           int32                                          `json:"toPeriod,omitempty"`
	AnnualInterestRate float64                                        `json:"annualInterestRate,omitempty"`
	Currency           GetRecurringDepositProductsProductIdCurrency   `json:"currency,omitempty"`
}
