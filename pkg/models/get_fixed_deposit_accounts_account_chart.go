package models

// GetFixedDepositAccountsAccountChart struct for GetFixedDepositAccountsAccountChart
type GetFixedDepositAccountsAccountChart struct {
	Id            int32                                `json:"id,omitempty"`
	FromDate      string                               `json:"fromDate,omitempty"`
	AccountId     int32                                `json:"accountId,omitempty"`
	AccountNumber int64                                `json:"accountNumber,omitempty"`
	ChartSlabs    []GetFixedDepositAccountsChartSlabs  `json:"chartSlabs,omitempty"`
	PeriodTypes   []GetFixedDepositAccountsPeriodTypes `json:"periodTypes,omitempty"`
}
