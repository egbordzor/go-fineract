package models

// OfficeData struct for OfficeData
type OfficeData struct {
	Id          int64  `json:"id,omitempty"`
	OpeningDate string `json:"openingDate,omitempty"`
	Hierarchy   string `json:"hierarchy,omitempty"`
	RowIndex    int32  `json:"rowIndex,omitempty"`
}
