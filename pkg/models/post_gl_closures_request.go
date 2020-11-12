package models

// PostGlClosuresRequest PostGLCLosuresRequest
type PostGlClosuresRequest struct {
	OfficeId    int64  `json:"officeId,omitempty"`
	ClosingDate string `json:"closingDate,omitempty"`
	Comments    string `json:"comments,omitempty"`
	Locale      string `json:"locale,omitempty"`
	DateFormat  string `json:"dateFormat,omitempty"`
}
