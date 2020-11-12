package models

// IdDocument struct for IdDocument
type IdDocument struct {
	IdType             string `json:"idType,omitempty"`
	IdNumber           string `json:"idNumber,omitempty"`
	IssuerCountry      string `json:"issuerCountry,omitempty"`
	OtherIdDescription string `json:"otherIdDescription,omitempty"`
}
