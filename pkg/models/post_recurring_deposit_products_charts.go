package models

// PostRecurringDepositProductsCharts struct for PostRecurringDepositProductsCharts
type PostRecurringDepositProductsCharts struct {
	FromDate   string                                   `json:"fromDate,omitempty"`
	Locale     string                                   `json:"locale,omitempty"`
	DateFormat string                                   `json:"dateFormat,omitempty"`
	ChartSlabs []PostRecurringDepositProductsChartSlabs `json:"chartSlabs,omitempty"`
}
