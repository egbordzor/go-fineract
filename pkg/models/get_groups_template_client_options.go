package models

// GetGroupsTemplateClientOptions struct for GetGroupsTemplateClientOptions
type GetGroupsTemplateClientOptions struct {
	Id          int32  `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	OfficeId    int32  `json:"officeId,omitempty"`
	OfficeName  string `json:"officeName,omitempty"`
}
