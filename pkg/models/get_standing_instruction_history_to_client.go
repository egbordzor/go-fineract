package models

// GetStandingInstructionHistoryToClient struct for GetStandingInstructionHistoryToClient
type GetStandingInstructionHistoryToClient struct {
	Id          int64  `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	OfficeId    int64  `json:"officeId,omitempty"`
	OfficeName  string `json:"officeName,omitempty"`
}
