package models

// PostClientClientIdAddressesRequest PostClientClientIdAddressesRequest
type PostClientClientIdAddressesRequest struct {
	Street          string `json:"street,omitempty"`
	AddressLine1    string `json:"addressLine1,omitempty"`
	AddressLine2    string `json:"addressLine2,omitempty"`
	AddressLine3    string `json:"addressLine3,omitempty"`
	City            string `json:"city,omitempty"`
	StateProvinceId int32  `json:"stateProvinceId,omitempty"`
	CountryId       int32  `json:"countryId,omitempty"`
	PostalCode      int64  `json:"postalCode,omitempty"`
}
