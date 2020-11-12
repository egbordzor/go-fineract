package models

// GetClientsStaffOptions struct for GetClientsStaffOptions
type GetClientsStaffOptions struct {
	Id            int32  `json:"id,omitempty"`
	Firstname     string `json:"firstname,omitempty"`
	Lastname      string `json:"lastname,omitempty"`
	DisplayName   string `json:"displayName,omitempty"`
	OfficeId      int32  `json:"officeId,omitempty"`
	OfficeName    string `json:"officeName,omitempty"`
	IsLoanOfficer bool   `json:"isLoanOfficer,omitempty"`
	IsActive      bool   `json:"isActive,omitempty"`
}
