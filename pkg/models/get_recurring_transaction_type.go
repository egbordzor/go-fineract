package models

// GetRecurringTransactionType struct for GetRecurringTransactionType
type GetRecurringTransactionType struct {
	Id                int32  `json:"id,omitempty"`
	Code              string `json:"code,omitempty"`
	Description       string `json:"description,omitempty"`
	Deposit           bool   `json:"deposit,omitempty"`
	Withdrawal        bool   `json:"withdrawal,omitempty"`
	InterestPosting   bool   `json:"interestPosting,omitempty"`
	FeeDeduction      bool   `json:"feeDeduction,omitempty"`
	InitiateTransfer  bool   `json:"initiateTransfer,omitempty"`
	ApproveTransfer   bool   `json:"approveTransfer,omitempty"`
	WithdrawTransfer  bool   `json:"withdrawTransfer,omitempty"`
	RejectTransfer    bool   `json:"rejectTransfer,omitempty"`
	OverdraftInterest bool   `json:"overdraftInterest,omitempty"`
	Writtenoff        bool   `json:"writtenoff,omitempty"`
	OverdraftFee      bool   `json:"overdraftFee,omitempty"`
}
