package models

// DeleteEntityTypeEntityIdDocumentsResponse DeleteEntityTypeEntityIdDocumentsResponse
type DeleteEntityTypeEntityIdDocumentsResponse struct {
	ResourceId         int64                  `json:"resourceId,omitempty"`
	Changes            map[string]interface{} `json:"changes,omitempty"`
	ResourceIdentifier string                 `json:"resourceIdentifier,omitempty"`
}
