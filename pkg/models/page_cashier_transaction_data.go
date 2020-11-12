package models

// PageCashierTransactionData struct for PageCashierTransactionData
type PageCashierTransactionData struct {
	TotalFilteredRecords int32                    `json:"totalFilteredRecords,omitempty"`
	PageItems            []CashierTransactionData `json:"pageItems,omitempty"`
}
