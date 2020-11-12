package models

// PutDataTablesAppTableIdDatatableIdRequest PutDataTablesAppTableIdDatatableIdRequest
type PutDataTablesAppTableIdDatatableIdRequest struct {
	DateOfBirth        string `json:"DateOfBirth,omitempty"`
	EducationCdHighest int64  `json:"Education_cdHighest,omitempty"`
	Name               string `json:"Name,omitempty"`
	OtherNotes         string `json:"OtherNotes,omitempty"`
	PointsScore        int64  `json:"PointsScore,omitempty"`
	DateFormat         string `json:"dateFormat,omitempty"`
	Locale             string `json:"locale,omitempty"`
}
