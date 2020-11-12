package models

// PutRecurringDepositProductsResponse PutRecurringDepositProductsResponse
type PutRecurringDepositProductsResponse struct {
	ResourceId int32                              `json:"resourceId,omitempty"`
	Changes    PutRecurringDepositProductsChanges `json:"changes,omitempty"`
}
