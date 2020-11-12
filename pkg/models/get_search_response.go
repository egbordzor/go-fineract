package models

// GetSearchResponse GetSearchResponse
type GetSearchResponse struct {
	EntityId         int64          `json:"entityId,omitempty"`
	EntityAccountNo  int64          `json:"entityAccountNo,omitempty"`
	EntityExternalId string         `json:"entityExternalId,omitempty"`
	EntityName       string         `json:"entityName,omitempty"`
	EntityType       string         `json:"entityType,omitempty"`
	ParentId         int64          `json:"parentId,omitempty"`
	ParentName       string         `json:"parentName,omitempty"`
	EntityStatus     EnumOptionData `json:"entityStatus,omitempty"`
}
