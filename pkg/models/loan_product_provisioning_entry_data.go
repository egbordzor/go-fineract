package models

// LoanProductProvisioningEntryData struct for LoanProductProvisioningEntryData
type LoanProductProvisioningEntryData struct {
	HistoryId          int64   `json:"historyId,omitempty"`
	OfficeId           int64   `json:"officeId,omitempty"`
	CurrencyCode       string  `json:"currencyCode,omitempty"`
	ProductId          int64   `json:"productId,omitempty"`
	CategoryId         int64   `json:"categoryId,omitempty"`
	OverdueInDays      int64   `json:"overdueInDays,omitempty"`
	Percentage         float32 `json:"percentage,omitempty"`
	LiablityAccount    int64   `json:"liablityAccount,omitempty"`
	ExpenseAccount     int64   `json:"expenseAccount,omitempty"`
	CriteriaId         int64   `json:"criteriaId,omitempty"`
	OutstandingBalance float32 `json:"outstandingBalance,omitempty"`
}
