package models

// PostGroupsGroupIdRequest PostGroupsGroupIdRequest
type PostGroupsGroupIdRequest struct {
	DestinationGroupId int32                      `json:"destinationGroupId,omitempty"`
	Clients            []PostGroupsGroupIdClients `json:"clients,omitempty"`
}
