package models

// GetAccountTransfersFromClientOptions struct for GetAccountTransfersFromClientOptions
type GetAccountTransfersFromClientOptions struct {
	Id          int32  `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	OfficeId    int32  `json:"officeId,omitempty"`
	OfficeName  string `json:"officeName,omitempty"`
}
