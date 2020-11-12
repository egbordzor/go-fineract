package models

// GetRecurringDepositProductsProductIdActiveChart struct for GetRecurringDepositProductsProductIdActiveChart
type GetRecurringDepositProductsProductIdActiveChart struct {
	Id                 int32                                            `json:"id,omitempty"`
	FromDate           string                                           `json:"fromDate,omitempty"`
	SavingsProductId   int32                                            `json:"savingsProductId,omitempty"`
	SavingsProductName string                                           `json:"savingsProductName,omitempty"`
	ChartSlabs         []GetRecurringDepositProductsProductIdChartSlabs `json:"chartSlabs,omitempty"`
	PeriodTypes        []GetRecurringDepositProductsProductIdPeriodType `json:"periodTypes,omitempty"`
}
