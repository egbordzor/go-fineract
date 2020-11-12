package models

// PostInterestRateChartsChartIdChartSlabsRequest PostInterestRateChartsChartIdChartSlabsRequest
type PostInterestRateChartsChartIdChartSlabsRequest struct {
	PeriodType         int32                                               `json:"periodType,omitempty"`
	FromPeriod         int32                                               `json:"fromPeriod,omitempty"`
	ToPeriod           int32                                               `json:"toPeriod,omitempty"`
	AnnualInterestRate float64                                             `json:"annualInterestRate,omitempty"`
	Description        string                                              `json:"description,omitempty"`
	Locale             string                                              `json:"locale,omitempty"`
	Incentives         []PostInterestRateChartsChartIdChartSlabsIncentives `json:"incentives,omitempty"`
}
