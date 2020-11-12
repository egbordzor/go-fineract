package models

// PutFixedDepositProductsProductIdResponse PutFixedDepositProductsProductIdResponse
type PutFixedDepositProductsProductIdResponse struct {
	ResourceId int32                          `json:"resourceId,omitempty"`
	Changes    PutFixedDepositProductsChanges `json:"changes,omitempty"`
}
