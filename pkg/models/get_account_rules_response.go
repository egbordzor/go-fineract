package models

// GetAccountRulesResponse GetAccountRulesResponse
type GetAccountRulesResponse struct {
	Id                         int64                   `json:"id,omitempty"`
	OfficeId                   int64                   `json:"officeId,omitempty"`
	OfficeName                 string                  `json:"officeName,omitempty"`
	Name                       string                  `json:"name,omitempty"`
	Description                string                  `json:"description,omitempty"`
	SystemDefined              bool                    `json:"systemDefined,omitempty"`
	AllowMultipleDebitEntries  bool                    `json:"allowMultipleDebitEntries,omitempty"`
	AllowMultipleCreditEntries bool                    `json:"allowMultipleCreditEntries,omitempty"`
	CreditTags                 []AccountingTagRuleData `json:"creditTags,omitempty"`
	DebitTags                  []AccountingTagRuleData `json:"debitTags,omitempty"`
}
