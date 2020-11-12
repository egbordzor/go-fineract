package models

// GetSelfClientsPageItems struct for GetSelfClientsPageItems
type GetSelfClientsPageItems struct {
	Id             int32                `json:"id,omitempty"`
	AccountNo      int64                `json:"accountNo,omitempty"`
	Status         GetSelfClientsStatus `json:"status,omitempty"`
	Active         bool                 `json:"active,omitempty"`
	ActivationDate string               `json:"activationDate,omitempty"`
	Fullname       string               `json:"fullname,omitempty"`
	DisplayName    string               `json:"displayName,omitempty"`
	OfficeId       int32                `json:"officeId,omitempty"`
	OfficeName     string               `json:"officeName,omitempty"`
}
