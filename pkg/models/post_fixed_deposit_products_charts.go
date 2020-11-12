package models

// PostFixedDepositProductsCharts struct for PostFixedDepositProductsCharts
type PostFixedDepositProductsCharts struct {
	FromDate   string                               `json:"fromDate,omitempty"`
	Locale     string                               `json:"locale,omitempty"`
	DateFormat string                               `json:"dateFormat,omitempty"`
	ChartSlabs []PostFixedDepositProductsChartSlabs `json:"chartSlabs,omitempty"`
}
