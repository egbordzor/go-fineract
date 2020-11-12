package models

// PostSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest PostSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest
type PostSavingsAccountsSavingsAccountIdChargesSavingsAccountChargeIdRequest struct {
	DateFormat string  `json:"dateFormat,omitempty"`
	Locale     string  `json:"locale,omitempty"`
	Amount     float32 `json:"amount,omitempty"`
	DueDate    string  `json:"dueDate,omitempty"`
}
