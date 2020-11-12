package models

// GetTellersTellerIdCashiersCashiersIdTransactionsTemplateResponse GetTellersTellerIdCashiersCashiersIdTransactionsTemplateResponse
type GetTellersTellerIdCashiersCashiersIdTransactionsTemplateResponse struct {
	CashierId       int64          `json:"cashierId,omitempty"`
	OfficeName      string         `json:"officeName,omitempty"`
	TellerId        int64          `json:"tellerId,omitempty"`
	TellerName      string         `json:"tellerName,omitempty"`
	CashierName     string         `json:"cashierName,omitempty"`
	CashierData     CashierData    `json:"cashierData,omitempty"`
	StartDate       string         `json:"startDate,omitempty"`
	EndDate         string         `json:"endDate,omitempty"`
	CurrencyOptions []CurrencyData `json:"currencyOptions,omitempty"`
}
