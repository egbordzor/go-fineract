package models

import (
	"time"
)

// PostTellersTellerIdCashiersCashierIdAllocateRequest PostTellersTellerIdCashiersCashierIdAllocateRequest
type PostTellersTellerIdCashiersCashierIdAllocateRequest struct {
	CurrencyCode string    `json:"currencyCode,omitempty"`
	TxnAmount    float32   `json:"txnAmount,omitempty"`
	TxnNote      string    `json:"txnNote,omitempty"`
	Locale       string    `json:"locale,omitempty"`
	DateFormat   string    `json:"dateFormat,omitempty"`
	TxnDate      time.Time `json:"txnDate,omitempty"`
}
