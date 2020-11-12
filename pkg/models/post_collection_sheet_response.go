package models

// PostCollectionSheetResponse PostCollectionSheetResponse
type PostCollectionSheetResponse struct {
	GroupId    int32                      `json:"groupId,omitempty"`
	ResourceId int32                      `json:"resourceId,omitempty"`
	Changes    PostCollectionSheetChanges `json:"changes,omitempty"`
}
