package models

// PostOfficesRequest PostOfficesRequest
type PostOfficesRequest struct {
	Name        string `json:"name,omitempty"`
	DateFormat  string `json:"dateFormat,omitempty"`
	Locale      string `json:"locale,omitempty"`
	OpeningDate string `json:"openingDate,omitempty"`
	ParentId    int64  `json:"parentId,omitempty"`
	ExternalId  string `json:"externalId,omitempty"`
}
