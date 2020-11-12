package models

// Office struct for Office
type Office struct {
	Id               int64  `json:"id,omitempty"`
	Parent           Office `json:"parent,omitempty"`
	Name             string `json:"name,omitempty"`
	Hierarchy        string `json:"hierarchy,omitempty"`
	OpeningLocalDate string `json:"openingLocalDate,omitempty"`
	New              bool   `json:"new,omitempty"`
}
