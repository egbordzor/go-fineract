package models

import (
	"time"
)

// GetTellersTellerIdCashiersCashiersIdTransactionsResponse GetTellersTellerIdCashiersCashiersIdTransactionsResponse
type GetTellersTellerIdCashiersCashiersIdTransactionsResponse struct {
	Id          int64          `json:"id,omitempty"`
	CashierId   int64          `json:"cashierId,omitempty"`
	TxnType     CashierTxnType `json:"txnType,omitempty"`
	TxnAmount   float32        `json:"txnAmount,omitempty"`
	TxnDate     time.Time      `json:"txnDate,omitempty"`
	EntityId    int64          `json:"entityId,omitempty"`
	EntityType  string         `json:"entityType,omitempty"`
	TxnNote     string         `json:"txnNote,omitempty"`
	CreatedDate time.Time      `json:"createdDate,omitempty"`
	OfficeId    int64          `json:"officeId,omitempty"`
	OfficeName  string         `json:"officeName,omitempty"`
	TellerId    int64          `json:"tellerId,omitempty"`
	CashierName string         `json:"cashierName,omitempty"`
}
