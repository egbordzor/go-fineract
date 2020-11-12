package models

import (
	"time"
)

// FloatingRate struct for FloatingRate
type FloatingRate struct {
	Id                  int64                `json:"id,omitempty"`
	Name                string               `json:"name,omitempty"`
	FloatingRatePeriods []FloatingRatePeriod `json:"floatingRatePeriods,omitempty"`
	CreatedBy           AppUser              `json:"createdBy,omitempty"`
	ModifiedBy          AppUser              `json:"modifiedBy,omitempty"`
	CreatedOn           time.Time            `json:"createdOn,omitempty"`
	ModifiedOn          time.Time            `json:"modifiedOn,omitempty"`
	BaseLendingRate     bool                 `json:"baseLendingRate,omitempty"`
	Active              bool                 `json:"active,omitempty"`
	New                 bool                 `json:"new,omitempty"`
}
