package models

import (
	"time"
)

// TaxComponent struct for TaxComponent
type TaxComponent struct {
	Id                    int64                 `json:"id,omitempty"`
	CreatedBy             AppUser               `json:"createdBy,omitempty"`
	CreatedDate           time.Time             `json:"createdDate,omitempty"`
	LastModifiedBy        AppUser               `json:"lastModifiedBy,omitempty"`
	LastModifiedDate      time.Time             `json:"lastModifiedDate,omitempty"`
	Percentage            float32               `json:"percentage,omitempty"`
	DebitAccountType      int32                 `json:"debitAccountType,omitempty"`
	DebitAcount           GlAccount             `json:"debitAcount,omitempty"`
	CreditAccountType     int32                 `json:"creditAccountType,omitempty"`
	CreditAcount          GlAccount             `json:"creditAcount,omitempty"`
	TaxComponentHistories []TaxComponentHistory `json:"taxComponentHistories,omitempty"`
	TaxGroupMappings      []TaxGroupMappings    `json:"taxGroupMappings,omitempty"`
	New                   bool                  `json:"new,omitempty"`
}
