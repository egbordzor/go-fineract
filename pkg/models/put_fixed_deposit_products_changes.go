package models

// PutFixedDepositProductsChanges struct for PutFixedDepositProductsChanges
type PutFixedDepositProductsChanges struct {
	Description    string `json:"description,omitempty"`
	MinDepositTerm int32  `json:"minDepositTerm,omitempty"`
}
