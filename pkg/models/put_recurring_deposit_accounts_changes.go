package models

// PutRecurringDepositAccountsChanges struct for PutRecurringDepositAccountsChanges
type PutRecurringDepositAccountsChanges struct {
	DepositAmount int32  `json:"depositAmount,omitempty"`
	Locale        string `json:"locale,omitempty"`
}
