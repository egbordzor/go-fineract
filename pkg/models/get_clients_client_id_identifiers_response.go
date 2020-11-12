package models

// GetClientsClientIdIdentifiersResponse GetClientsClientIdIdentifiersResponse
type GetClientsClientIdIdentifiersResponse struct {
	Id           int32                  `json:"id,omitempty"`
	ClientId     int32                  `json:"clientId,omitempty"`
	DocumentType GetClientsDocumentType `json:"documentType,omitempty"`
	DocumentKey  string                 `json:"documentKey,omitempty"`
	Description  string                 `json:"description,omitempty"`
}
