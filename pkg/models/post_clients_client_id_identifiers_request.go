package models

// PostClientsClientIdIdentifiersRequest PostClientsClientIdIdentifiersRequest
type PostClientsClientIdIdentifiersRequest struct {
	DocumentTypeId int32  `json:"documentTypeId,omitempty"`
	DocumentKey    string `json:"documentKey,omitempty"`
	Description    string `json:"description,omitempty"`
}
