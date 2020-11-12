package models

// GetClientClientIdAddressesResponse GetClientClientIdAddressesResponse
type GetClientClientIdAddressesResponse struct {
	ClientId        int64  `json:"client_id,omitempty"`
	AddressType     string `json:"addressType,omitempty"`
	AddressId       int32  `json:"addressId,omitempty"`
	AddressTypeId   int32  `json:"addressTypeId,omitempty"`
	IsActive        bool   `json:"isActive,omitempty"`
	Street          string `json:"street,omitempty"`
	AddressLine1    string `json:"addressLine1,omitempty"`
	AddressLine2    string `json:"addressLine2,omitempty"`
	AddressLine3    string `json:"addressLine3,omitempty"`
	TownVillage     string `json:"townVillage,omitempty"`
	City            string `json:"city,omitempty"`
	CountyDistrict  string `json:"countyDistrict,omitempty"`
	StateProvinceId int32  `json:"stateProvinceId,omitempty"`
	CountryName     string `json:"countryName,omitempty"`
	StateName       string `json:"stateName,omitempty"`
	CountryId       int32  `json:"countryId,omitempty"`
	PostalCode      int64  `json:"postalCode,omitempty"`
	CreatedBy       string `json:"createdBy,omitempty"`
	UpdatedBy       string `json:"updatedBy,omitempty"`
}
