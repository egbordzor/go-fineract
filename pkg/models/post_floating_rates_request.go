package models

// PostFloatingRatesRequest PostFloatingRatesRequest
type PostFloatingRatesRequest struct {
	Name              string                         `json:"name,omitempty"`
	IsBaseLendingRate bool                           `json:"isBaseLendingRate,omitempty"`
	IsActive          bool                           `json:"isActive,omitempty"`
	RatePeriods       []PostFloatingRatesRatePeriods `json:"ratePeriods,omitempty"`
}
