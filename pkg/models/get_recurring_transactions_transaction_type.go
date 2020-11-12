package models

// GetRecurringTransactionsTransactionType struct for GetRecurringTransactionsTransactionType
type GetRecurringTransactionsTransactionType struct {
	Id              int32  `json:"id,omitempty"`
	Code            string `json:"code,omitempty"`
	Description     string `json:"description,omitempty"`
	Deposit         bool   `json:"deposit,omitempty"`
	Withdrawal      bool   `json:"withdrawal,omitempty"`
	InterestPosting bool   `json:"interestPosting,omitempty"`
	FeeDeduction    bool   `json:"feeDeduction,omitempty"`
}
