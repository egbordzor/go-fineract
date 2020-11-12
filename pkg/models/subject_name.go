package models

// SubjectName struct for SubjectName
type SubjectName struct {
	FirstName   string `json:"firstName,omitempty"`
	MiddleName  string `json:"middleName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}
