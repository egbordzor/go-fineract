package models

// PostRunaccrualsRequest runaccrualsRequest
type PostRunaccrualsRequest struct {
	Locale     string `json:"locale,omitempty"`
	DateFormat string `json:"dateFormat,omitempty"`
	// which specifies periodic accruals should happen till the given Date
	TillDate string `json:"tillDate"`
}
