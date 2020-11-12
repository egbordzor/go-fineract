package models

import (
	"time"
)

// ProvisioningEntryData struct for ProvisioningEntryData
type ProvisioningEntryData struct {
	Id          int64                              `json:"id,omitempty"`
	CreatedDate time.Time                          `json:"createdDate,omitempty"`
	Entries     []LoanProductProvisioningEntryData `json:"entries,omitempty"`
}
