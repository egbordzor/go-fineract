package models

import (
	"time"
)

// FloatingRatePeriod struct for FloatingRatePeriod
type FloatingRatePeriod struct {
	Id                            int64        `json:"id,omitempty"`
	FromDate                      time.Time    `json:"fromDate,omitempty"`
	InterestRate                  float32      `json:"interestRate,omitempty"`
	CreatedBy                     AppUser      `json:"createdBy,omitempty"`
	ModifiedBy                    AppUser      `json:"modifiedBy,omitempty"`
	CreatedOn                     time.Time    `json:"createdOn,omitempty"`
	ModifiedOn                    time.Time    `json:"modifiedOn,omitempty"`
	DifferentialToBaseLendingRate bool         `json:"differentialToBaseLendingRate,omitempty"`
	FloatingRatesId               FloatingRate `json:"floatingRatesId,omitempty"`
	Active                        bool         `json:"active,omitempty"`
	New                           bool         `json:"new,omitempty"`
}
