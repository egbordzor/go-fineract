package models

// PutFloatingRatesChanges struct for PutFloatingRatesChanges
type PutFloatingRatesChanges struct {
	RatePeriods []PostFloatingRatesRatePeriods `json:"ratePeriods,omitempty"`
}
