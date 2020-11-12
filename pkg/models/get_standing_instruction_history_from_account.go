package models

// GetStandingInstructionHistoryFromAccount struct for GetStandingInstructionHistoryFromAccount
type GetStandingInstructionHistoryFromAccount struct {
	Id          int64  `json:"id,omitempty"`
	AccountNo   int64  `json:"accountNo,omitempty"`
	ProductId   int64  `json:"productId,omitempty"`
	ProductName string `json:"productName,omitempty"`
}
