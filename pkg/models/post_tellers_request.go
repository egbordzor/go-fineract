package models

// PostTellersRequest PostTellersRequest
type PostTellersRequest struct {
	OfficeId    int64  `json:"officeId,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	Locale      string `json:"locale,omitempty"`
	DateFormat  string `json:"dateFormat,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
}
