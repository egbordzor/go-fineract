package models

// PutEntityTypeEntityIdDocumentsResponse PutEntityTypeEntityIdDocumentsResponse
type PutEntityTypeEntityIdDocumentsResponse struct {
	ResourceId         int64                  `json:"resourceId,omitempty"`
	Changes            map[string]interface{} `json:"changes,omitempty"`
	ResourceIdentifier string                 `json:"resourceIdentifier,omitempty"`
}
