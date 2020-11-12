package models

// PutFixedDepositProductsProductIdRequest PutFixedDepositProductsProductIdRequest
type PutFixedDepositProductsProductIdRequest struct {
	Description          string `json:"description,omitempty"`
	Locale               string `json:"locale,omitempty"`
	MinDepositTerm       int32  `json:"minDepositTerm,omitempty"`
	MinDepositTermTypeId int32  `json:"minDepositTermTypeId,omitempty"`
}
