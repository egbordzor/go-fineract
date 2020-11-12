package models

// PostUsersRequest PostUsersRequest
type PostUsersRequest struct {
	Username            string `json:"username,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Email               string `json:"email,omitempty"`
	OfficeId            int64  `json:"officeId,omitempty"`
	StaffId             int64  `json:"staffId,omitempty"`
	Roles               string `json:"roles,omitempty"`
	SendPasswordToEmail bool   `json:"sendPasswordToEmail,omitempty"`
	IsSelfServiceUser   bool   `json:"isSelfServiceUser,omitempty"`
}
