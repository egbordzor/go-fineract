package models

// GetInterestRateChartsResponse GetInterestRateChartsResponse
type GetInterestRateChartsResponse struct {
	Id                 int32                             `json:"id,omitempty"`
	FromDate           string                            `json:"fromDate,omitempty"`
	SavingsProductId   int32                             `json:"savingsProductId,omitempty"`
	SavingsProductName string                            `json:"savingsProductName,omitempty"`
	ChartSlabs         []GetInterestRateChartsChartSlabs `json:"chartSlabs,omitempty"`
}
