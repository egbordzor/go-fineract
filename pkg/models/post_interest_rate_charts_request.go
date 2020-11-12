package models

// PostInterestRateChartsRequest PostInterestRateChartsRequest
type PostInterestRateChartsRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Locale      string `json:"locale,omitempty"`
	DateFormat  string `json:"dateFormat,omitempty"`
	FromDate    string `json:"fromDate,omitempty"`
}
