package models

import (
	"time"
)

// InteropTransferRequestData struct for InteropTransferRequestData
type InteropTransferRequestData struct {
	TransactionCode     string                     `json:"transactionCode"`
	RequestCode         string                     `json:"requestCode,omitempty"`
	AccountId           string                     `json:"accountId"`
	Amount              MoneyData                  `json:"amount"`
	TransactionRole     string                     `json:"transactionRole"`
	TransactionType     InteropTransactionTypeData `json:"transactionType,omitempty"`
	Note                string                     `json:"note,omitempty"`
	GeoCode             GeoCodeData                `json:"geoCode,omitempty"`
	Expiration          time.Time                  `json:"expiration,omitempty"`
	ExtensionList       []ExtensionData            `json:"extensionList,omitempty"`
	TransferCode        string                     `json:"transferCode"`
	FspFee              MoneyData                  `json:"fspFee,omitempty"`
	FspCommission       MoneyData                  `json:"fspCommission,omitempty"`
	ExpirationLocalDate string                     `json:"expirationLocalDate,omitempty"`
}
