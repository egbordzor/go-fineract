package models

// AccountingRuleData struct for AccountingRuleData
type AccountingRuleData struct {
	CreditTags []AccountingTagRuleData `json:"creditTags,omitempty"`
	DebitTags  []AccountingTagRuleData `json:"debitTags,omitempty"`
}
