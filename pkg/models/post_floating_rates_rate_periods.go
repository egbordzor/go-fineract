package models

// PostFloatingRatesRatePeriods struct for PostFloatingRatesRatePeriods
type PostFloatingRatesRatePeriods struct {
	FromDate     string  `json:"fromDate,omitempty"`
	InterestRate float64 `json:"interestRate,omitempty"`
	Locale       string  `json:"locale,omitempty"`
	DateFormat   string  `json:"dateFormat,omitempty"`
}
