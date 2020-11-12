package models

// PostStaffRequest PostStaffRequest
type PostStaffRequest struct {
	Id            int64  `json:"id,omitempty"`
	Firstname     string `json:"firstname,omitempty"`
	Lastname      string `json:"lastname,omitempty"`
	IsLoanOfficer bool   `json:"isLoanOfficer,omitempty"`
	ExternalId    string `json:"externalId,omitempty"`
	MobileNo      string `json:"mobileNo,omitempty"`
	IsActive      bool   `json:"isActive,omitempty"`
	JoiningDate   string `json:"joiningDate,omitempty"`
	Locale        string `json:"locale,omitempty"`
	DateFormat    string `json:"dateFormat,omitempty"`
}
