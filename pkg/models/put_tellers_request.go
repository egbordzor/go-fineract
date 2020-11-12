package models

// PutTellersRequest PutTellersRequest
type PutTellersRequest struct {
	Name        string `json:"name,omitempty"`
	OfficeId    int64  `json:"officeId,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	EndDate     string `json:"endDate,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
	Locale      string `json:"locale,omitempty"`
	DateFormat  string `json:"dateFormat,omitempty"`
}
