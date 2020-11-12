package models

// Grouping struct for Grouping
type Grouping struct {
	Name     string   `json:"name,omitempty"`
	Entities []Entity `json:"entities,omitempty"`
}
