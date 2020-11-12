package models

// PostHolidaysRequest PostHolidaysRequest
type PostHolidaysRequest struct {
	Name                    string                       `json:"name,omitempty"`
	Description             string                       `json:"description,omitempty"`
	DateFormat              string                       `json:"dateFormat,omitempty"`
	Locale                  string                       `json:"locale,omitempty"`
	FromDate                string                       `json:"fromDate,omitempty"`
	ToDate                  string                       `json:"toDate,omitempty"`
	RepaymentsRescheduledTo string                       `json:"repaymentsRescheduledTo,omitempty"`
	Offices                 []PostHolidaysRequestOffices `json:"offices,omitempty"`
}
