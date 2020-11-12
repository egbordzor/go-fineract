package models

// GetInterestRateChartsChartSlabs struct for GetInterestRateChartsChartSlabs
type GetInterestRateChartsChartSlabs struct {
	Id                 int32                                    `json:"id,omitempty"`
	PeriodTypes        GetInterestRateChartsTemplatePeriodTypes `json:"periodTypes,omitempty"`
	FromPeriod         int32                                    `json:"fromPeriod,omitempty"`
	AnnualInterestRate int32                                    `json:"annualInterestRate,omitempty"`
	Currency           GetInterestRateChartsCurrency            `json:"currency,omitempty"`
}
