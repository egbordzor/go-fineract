package models

// Staff struct for Staff
type Staff struct {
	Id             int64 `json:"id,omitempty"`
	LoanOfficer    bool  `json:"loanOfficer,omitempty"`
	Active         bool  `json:"active,omitempty"`
	Image          Image `json:"image,omitempty"`
	NotLoanOfficer bool  `json:"notLoanOfficer,omitempty"`
	NotActive      bool  `json:"notActive,omitempty"`
	New            bool  `json:"new,omitempty"`
}
