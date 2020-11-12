package models

// PostPaymentTypesRequest PostPaymentTypesRequest
type PostPaymentTypesRequest struct {
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	IsCashPayment bool   `json:"isCashPayment,omitempty"`
	Position      int32  `json:"position,omitempty"`
}
