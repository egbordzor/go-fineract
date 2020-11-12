package models

// PostAccountingRulesRequest PostAccountingRulesRequest
type PostAccountingRulesRequest struct {
	Name            string `json:"name,omitempty"`
	OfficeId        int64  `json:"officeId,omitempty"`
	AccountToDebit  int64  `json:"accountToDebit,omitempty"`
	AccountToCredit int64  `json:"accountToCredit,omitempty"`
	Description     string `json:"description,omitempty"`
}
