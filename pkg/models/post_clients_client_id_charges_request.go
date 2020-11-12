package models

// PostClientsClientIdChargesRequest PostClientsClientIdChargesRequest
type PostClientsClientIdChargesRequest struct {
	Amount     int32  `json:"amount,omitempty"`
	ChargeId   int32  `json:"chargeId,omitempty"`
	DateFormat string `json:"dateFormat,omitempty"`
	DueDate    string `json:"dueDate,omitempty"`
	Locale     string `json:"locale,omitempty"`
}
