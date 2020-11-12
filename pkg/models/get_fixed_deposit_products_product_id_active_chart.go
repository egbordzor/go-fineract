package models

// GetFixedDepositProductsProductIdActiveChart struct for GetFixedDepositProductsProductIdActiveChart
type GetFixedDepositProductsProductIdActiveChart struct {
	Id                 int32                                        `json:"id,omitempty"`
	FromDate           string                                       `json:"fromDate,omitempty"`
	SavingsProductId   int32                                        `json:"savingsProductId,omitempty"`
	SavingsProductName string                                       `json:"savingsProductName,omitempty"`
	ChartSlabs         []GetFixedDepositProductsProductIdChartSlabs `json:"chartSlabs,omitempty"`
	PeriodTypes        []GetFixedDepositProductsProductIdPeriodType `json:"periodTypes,omitempty"`
}
