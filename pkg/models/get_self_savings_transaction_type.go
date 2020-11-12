package models

// GetSelfSavingsTransactionType struct for GetSelfSavingsTransactionType
type GetSelfSavingsTransactionType struct {
	Id              int32  `json:"id,omitempty"`
	Code            string `json:"code,omitempty"`
	Description     string `json:"description,omitempty"`
	Deposit         bool   `json:"deposit,omitempty"`
	Withdrawal      bool   `json:"withdrawal,omitempty"`
	InterestPosting bool   `json:"interestPosting,omitempty"`
	FeeDeduction    bool   `json:"feeDeduction,omitempty"`
}
