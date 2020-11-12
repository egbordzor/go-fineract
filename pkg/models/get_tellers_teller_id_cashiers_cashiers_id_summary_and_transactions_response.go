package models

// GetTellersTellerIdCashiersCashiersIdSummaryAndTransactionsResponse GetTellersTellerIdCashiersCashiersIdSummaryAndTransactionsResponse
type GetTellersTellerIdCashiersCashiersIdSummaryAndTransactionsResponse struct {
	SumCashAllocation   float32                    `json:"sumCashAllocation,omitempty"`
	SumInwardCash       float32                    `json:"sumInwardCash,omitempty"`
	SumOutwardCash      float32                    `json:"sumOutwardCash,omitempty"`
	SumCashSettlement   float32                    `json:"sumCashSettlement,omitempty"`
	NetCash             float32                    `json:"netCash,omitempty"`
	OfficeName          string                     `json:"officeName,omitempty"`
	TellerId            int64                      `json:"tellerId,omitempty"`
	TellerName          string                     `json:"tellerName,omitempty"`
	CashierId           int64                      `json:"cashierId,omitempty"`
	CashierName         string                     `json:"cashierName,omitempty"`
	CashierTransactions PageCashierTransactionData `json:"cashierTransactions,omitempty"`
}
