package models

// GetGroupsResponse GetGroupsResponse
type GetGroupsResponse struct {
	TotalFilteredRecords int32                `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetGroupsPageItems `json:"pageItems,omitempty"`
}
