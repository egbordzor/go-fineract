package models

import (
	"time"
)

// TaxComponentHistory struct for TaxComponentHistory
type TaxComponentHistory struct {
	Id               int64     `json:"id,omitempty"`
	CreatedBy        AppUser   `json:"createdBy,omitempty"`
	CreatedDate      time.Time `json:"createdDate,omitempty"`
	LastModifiedBy   AppUser   `json:"lastModifiedBy,omitempty"`
	LastModifiedDate time.Time `json:"lastModifiedDate,omitempty"`
	Percentage       float32   `json:"percentage,omitempty"`
	New              bool      `json:"new,omitempty"`
}
