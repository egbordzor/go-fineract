package models

// InteropIdentifierRequestData struct for InteropIdentifierRequestData
type InteropIdentifierRequestData struct {
	IdType      string `json:"idType"`
	IdValue     string `json:"idValue"`
	SubIdOrType string `json:"subIdOrType,omitempty"`
	AccountId   string `json:"accountId"`
}
