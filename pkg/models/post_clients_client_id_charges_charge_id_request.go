package models

// PostClientsClientIdChargesChargeIdRequest PostClientsClientIdChargesChargeIdRequest
type PostClientsClientIdChargesChargeIdRequest struct {
	Amount          int32  `json:"amount,omitempty"`
	Locale          string `json:"locale,omitempty"`
	DateFormat      string `json:"dateFormat,omitempty"`
	TransactionDate string `json:"transactionDate,omitempty"`
}
