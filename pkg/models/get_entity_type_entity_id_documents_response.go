package models

// GetEntityTypeEntityIdDocumentsResponse GetEntityTypeEntityIdDocumentsResponse
type GetEntityTypeEntityIdDocumentsResponse struct {
	Id               int64  `json:"id,omitempty"`
	ParentEntityType string `json:"parentEntityType,omitempty"`
	ParentEntityId   int64  `json:"parentEntityId,omitempty"`
	Name             string `json:"name,omitempty"`
	FileName         string `json:"fileName,omitempty"`
	Size             int64  `json:"size,omitempty"`
	Type             string `json:"type,omitempty"`
	Description      string `json:"description,omitempty"`
	Location         string `json:"location,omitempty"`
	StorageType      int32  `json:"storageType,omitempty"`
}
