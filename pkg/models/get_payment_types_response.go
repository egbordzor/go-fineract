package models

// GetPaymentTypesResponse GetPaymentTypesResponse
type GetPaymentTypesResponse struct {
	Id            int32  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	IsCashPayment bool   `json:"isCashPayment,omitempty"`
	Position      int32  `json:"position,omitempty"`
}
