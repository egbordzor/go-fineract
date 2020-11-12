package models

// Event struct for Event
type Event struct {
	ActionName string `json:"actionName,omitempty"`
	EntityName string `json:"entityName,omitempty"`
}
