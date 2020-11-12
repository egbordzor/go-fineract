package models

// PutRecurringDepositProductsChanges struct for PutRecurringDepositProductsChanges
type PutRecurringDepositProductsChanges struct {
	Description    string `json:"description,omitempty"`
	MinDepositTerm int32  `json:"minDepositTerm,omitempty"`
}
