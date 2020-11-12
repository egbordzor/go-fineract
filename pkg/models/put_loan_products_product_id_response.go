package models

// PutLoanProductsProductIdResponse PutLoanProductsProductIdResponse
type PutLoanProductsProductIdResponse struct {
	ResourceId int32          `json:"resourceId,omitempty"`
	Changes    PutLoanChanges `json:"changes,omitempty"`
}
