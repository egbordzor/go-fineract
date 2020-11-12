package models

// PostChargesRequest PostChargesRequest
type PostChargesRequest struct {
	Name                  string  `json:"name,omitempty"`
	ChargeAppliesTo       int32   `json:"chargeAppliesTo,omitempty"`
	CurrencyCode          string  `json:"currencyCode,omitempty"`
	Locale                string  `json:"locale,omitempty"`
	Amount                float32 `json:"amount,omitempty"`
	ChargeTimeType        int32   `json:"chargeTimeType,omitempty"`
	ChargeCalculationType int32   `json:"chargeCalculationType,omitempty"`
	ChargePaymentMode     int32   `json:"chargePaymentMode,omitempty"`
	Active                string  `json:"active,omitempty"`
}
