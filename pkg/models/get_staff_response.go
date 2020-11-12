package models

// GetStaffResponse GetStaffResponse
type GetStaffResponse struct {
	Id            int64  `json:"id,omitempty"`
	Firstname     string `json:"firstname,omitempty"`
	Lastname      string `json:"lastname,omitempty"`
	DisplayName   string `json:"displayName,omitempty"`
	OfficeId      int64  `json:"officeId,omitempty"`
	OfficeName    string `json:"officeName,omitempty"`
	IsLoanOfficer bool   `json:"isLoanOfficer,omitempty"`
	ExternalId    string `json:"externalId,omitempty"`
	IsActive      bool   `json:"isActive,omitempty"`
	JoiningDate   string `json:"joiningDate,omitempty"`
}
