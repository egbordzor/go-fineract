package models

// InteropAccountData struct for InteropAccountData
type InteropAccountData struct {
	OfficeId            int64                             `json:"officeId,omitempty"`
	GroupId             int64                             `json:"groupId,omitempty"`
	ClientId            int64                             `json:"clientId,omitempty"`
	LoanId              int64                             `json:"loanId,omitempty"`
	SavingsId           int64                             `json:"savingsId,omitempty"`
	SubResourceId       int64                             `json:"subResourceId,omitempty"`
	TransactionId       string                            `json:"transactionId,omitempty"`
	Changes             map[string]map[string]interface{} `json:"changes,omitempty"`
	ProductId           int64                             `json:"productId,omitempty"`
	GsimId              int64                             `json:"gsimId,omitempty"`
	GlimId              int64                             `json:"glimId,omitempty"`
	RollbackTransaction bool                              `json:"rollbackTransaction,omitempty"`
}
