package models

// PostDataTablesAppTableIdRequest PostDataTablesAppTableIdRequest
type PostDataTablesAppTableIdRequest struct {
	BusinessDescription string  `json:"BusinessDescription,omitempty"`
	Comment             string  `json:"Comment,omitempty"`
	EducationCv         string  `json:"Education_cv,omitempty"`
	GenderCd            int64   `json:"Gender_cd,omitempty"`
	HighestRatePaid     float64 `json:"HighestRatePaid,omitempty"`
	NextVisit           string  `json:"NextVisit,omitempty"`
	YearsinBusiness     int64   `json:"YearsinBusiness,omitempty"`
	DateFormat          string  `json:"dateFormat,omitempty"`
	Locale              string  `json:"locale,omitempty"`
}
