package models

// StaffData struct for StaffData
type StaffData struct {
	Id          int64  `json:"id,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	OfficeId    int64  `json:"officeId,omitempty"`
	OfficeName  string `json:"officeName,omitempty"`
	JoiningDate string `json:"joiningDate,omitempty"`
	RowIndex    int32  `json:"rowIndex,omitempty"`
}
