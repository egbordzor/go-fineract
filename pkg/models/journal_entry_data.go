package models

// JournalEntryData struct for JournalEntryData
type JournalEntryData struct {
	Id              int64          `json:"id,omitempty"`
	OfficeId        int64          `json:"officeId,omitempty"`
	GlAccountId     int64          `json:"glAccountId,omitempty"`
	GlAccountType   EnumOptionData `json:"glAccountType,omitempty"`
	TransactionDate string         `json:"transactionDate,omitempty"`
	EntryType       EnumOptionData `json:"entryType,omitempty"`
	Amount          float32        `json:"amount,omitempty"`
	TransactionId   string         `json:"transactionId,omitempty"`
	RowIndex        int32          `json:"rowIndex,omitempty"`
}
