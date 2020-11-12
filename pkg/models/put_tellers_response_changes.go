package models

// PutTellersResponseChanges struct for PutTellersResponseChanges
type PutTellersResponseChanges struct {
	Description string `json:"description,omitempty"`
	EndDate     string `json:"endDate,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
	Locale      string `json:"locale,omitempty"`
	DateFormat  string `json:"dateFormat,omitempty"`
}
