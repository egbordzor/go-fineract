package models

// JournalEntryCommand struct for JournalEntryCommand
type JournalEntryCommand struct {
	OfficeId         int64                             `json:"officeId,omitempty"`
	TransactionDate  string                            `json:"transactionDate,omitempty"`
	Comments         string                            `json:"comments,omitempty"`
	ReferenceNumber  string                            `json:"referenceNumber,omitempty"`
	AccountingRuleId int64                             `json:"accountingRuleId,omitempty"`
	Credits          []SingleDebitOrCreditEntryCommand `json:"credits,omitempty"`
	Debits           []SingleDebitOrCreditEntryCommand `json:"debits,omitempty"`
}
