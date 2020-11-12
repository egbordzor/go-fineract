package models

// PostalAddress struct for PostalAddress
type PostalAddress struct {
	AddressLine1  string `json:"addressLine1,omitempty"`
	AddressLine2  string `json:"addressLine2,omitempty"`
	City          string `json:"city,omitempty"`
	StateProvince string `json:"stateProvince,omitempty"`
	PostalCode    string `json:"postalCode,omitempty"`
	Country       string `json:"country,omitempty"`
}
