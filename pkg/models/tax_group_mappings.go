package models

import (
	"time"
)

// TaxGroupMappings struct for TaxGroupMappings
type TaxGroupMappings struct {
	Id               int64        `json:"id,omitempty"`
	CreatedBy        AppUser      `json:"createdBy,omitempty"`
	CreatedDate      time.Time    `json:"createdDate,omitempty"`
	LastModifiedBy   AppUser      `json:"lastModifiedBy,omitempty"`
	LastModifiedDate time.Time    `json:"lastModifiedDate,omitempty"`
	TaxComponent     TaxComponent `json:"taxComponent,omitempty"`
	EndDate          time.Time    `json:"endDate,omitempty"`
	New              bool         `json:"new,omitempty"`
}
