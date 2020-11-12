package models

import (
	"time"
)

// TaxGroup struct for TaxGroup
type TaxGroup struct {
	Id               int64              `json:"id,omitempty"`
	CreatedBy        AppUser            `json:"createdBy,omitempty"`
	CreatedDate      time.Time          `json:"createdDate,omitempty"`
	LastModifiedBy   AppUser            `json:"lastModifiedBy,omitempty"`
	LastModifiedDate time.Time          `json:"lastModifiedDate,omitempty"`
	Name             string             `json:"name,omitempty"`
	TaxGroupMappings []TaxGroupMappings `json:"taxGroupMappings,omitempty"`
	New              bool               `json:"new,omitempty"`
}
