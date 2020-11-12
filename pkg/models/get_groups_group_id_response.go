package models

// GetGroupsGroupIdResponse GetGroupsGroupIdResponse
type GetGroupsGroupIdResponse struct {
	Id         int32                    `json:"id,omitempty"`
	Name       string                   `json:"name,omitempty"`
	ExternalId string                   `json:"externalId,omitempty"`
	OfficeId   int32                    `json:"officeId,omitempty"`
	OfficeName string                   `json:"officeName,omitempty"`
	Hierarchy  string                   `json:"hierarchy,omitempty"`
	Timeline   GetGroupsGroupIdTimeline `json:"timeline,omitempty"`
}
