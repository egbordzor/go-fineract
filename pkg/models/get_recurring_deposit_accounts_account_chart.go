package models

// GetRecurringDepositAccountsAccountChart struct for GetRecurringDepositAccountsAccountChart
type GetRecurringDepositAccountsAccountChart struct {
	Id            int32                                    `json:"id,omitempty"`
	FromDate      string                                   `json:"fromDate,omitempty"`
	AccountId     int32                                    `json:"accountId,omitempty"`
	AccountNumber int64                                    `json:"accountNumber,omitempty"`
	ChartSlabs    []GetRecurringDepositAccountsChartSlabs  `json:"chartSlabs,omitempty"`
	PeriodTypes   []GetRecurringDepositAccountsPeriodTypes `json:"periodTypes,omitempty"`
}
