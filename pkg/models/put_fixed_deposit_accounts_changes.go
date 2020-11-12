package models

// PutFixedDepositAccountsChanges struct for PutFixedDepositAccountsChanges
type PutFixedDepositAccountsChanges struct {
	DepositAmount float32 `json:"depositAmount,omitempty"`
	Locale        string  `json:"locale,omitempty"`
}
