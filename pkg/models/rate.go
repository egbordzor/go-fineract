package models

import (
	"time"
)

// Rate struct for Rate
type Rate struct {
	Id               int64     `json:"id,omitempty"`
	CreatedBy        AppUser   `json:"createdBy,omitempty"`
	CreatedDate      time.Time `json:"createdDate,omitempty"`
	LastModifiedBy   AppUser   `json:"lastModifiedBy,omitempty"`
	LastModifiedDate time.Time `json:"lastModifiedDate,omitempty"`
	Name             string    `json:"name,omitempty"`
	Percentage       float32   `json:"percentage,omitempty"`
	ProductApply     int32     `json:"productApply,omitempty"`
	Active           bool      `json:"active,omitempty"`
	ApproveUser      AppUser   `json:"approveUser,omitempty"`
	New              bool      `json:"new,omitempty"`
}
