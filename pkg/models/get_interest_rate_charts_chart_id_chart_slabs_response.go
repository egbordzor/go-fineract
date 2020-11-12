package models

// GetInterestRateChartsChartIdChartSlabsResponse GetInterestRateChartsChartIdChartSlabsResponse
type GetInterestRateChartsChartIdChartSlabsResponse struct {
	Id                 int32                                              `json:"id,omitempty"`
	Description        string                                             `json:"description,omitempty"`
	PeriodTypes        GetInterestRateChartsTemplatePeriodTypes           `json:"periodTypes,omitempty"`
	FromPeriod         int32                                              `json:"fromPeriod,omitempty"`
	ToPeriod           int32                                              `json:"toPeriod,omitempty"`
	AnnualInterestRate float64                                            `json:"annualInterestRate,omitempty"`
	Currency           GetInterestRateChartsCurrency                      `json:"currency,omitempty"`
	Incentives         []GetInterestRateChartsChartIdChartSlabsIncentives `json:"incentives,omitempty"`
}
