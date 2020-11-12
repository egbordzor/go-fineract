package models

// GetClientsPageItemsResponse struct for GetClientsPageItemsResponse
type GetClientsPageItemsResponse struct {
	Id          int32           `json:"id,omitempty"`
	AccountNo   string          `json:"accountNo,omitempty"`
	Status      GetClientStatus `json:"status,omitempty"`
	Active      bool            `json:"active,omitempty"`
	Fullname    string          `json:"fullname,omitempty"`
	DisplayName string          `json:"displayName,omitempty"`
	OfficeId    int32           `json:"officeId,omitempty"`
	OfficeName  string          `json:"officeName,omitempty"`
}
