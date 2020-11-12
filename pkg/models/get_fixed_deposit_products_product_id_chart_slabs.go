package models

// GetFixedDepositProductsProductIdChartSlabs struct for GetFixedDepositProductsProductIdChartSlabs
type GetFixedDepositProductsProductIdChartSlabs struct {
	Id                 int32                                      `json:"id,omitempty"`
	Description        string                                     `json:"description,omitempty"`
	PeriodType         GetFixedDepositProductsProductIdPeriodType `json:"periodType,omitempty"`
	FromPeriod         int32                                      `json:"fromPeriod,omitempty"`
	ToPeriod           int32                                      `json:"toPeriod,omitempty"`
	AnnualInterestRate float64                                    `json:"annualInterestRate,omitempty"`
	Currency           GetFixedDepositProductsProductIdCurrency   `json:"currency,omitempty"`
}
