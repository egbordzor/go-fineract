package models

// GetSelfLoansLoanIdTransactionsType struct for GetSelfLoansLoanIdTransactionsType
type GetSelfLoansLoanIdTransactionsType struct {
	Id                      int32  `json:"id,omitempty"`
	Code                    string `json:"code,omitempty"`
	Description             string `json:"description,omitempty"`
	Disbursement            bool   `json:"disbursement,omitempty"`
	RepaymentAtDisbursement bool   `json:"repaymentAtDisbursement,omitempty"`
	Repayment               bool   `json:"repayment,omitempty"`
	Contra                  bool   `json:"contra,omitempty"`
	WaiveInterest           bool   `json:"waiveInterest,omitempty"`
	WaiveCharges            bool   `json:"waiveCharges,omitempty"`
	WriteOff                bool   `json:"writeOff,omitempty"`
	RecoveryRepayment       bool   `json:"recoveryRepayment,omitempty"`
}
