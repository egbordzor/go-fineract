package models

// PutSavingsProductsProductIdResponse PutSavingsProductsProductIdResponse
type PutSavingsProductsProductIdResponse struct {
	ResourceId int32             `json:"resourceId,omitempty"`
	Changes    PutSavingsChanges `json:"changes,omitempty"`
}
