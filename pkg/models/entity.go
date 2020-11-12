package models

// Entity struct for Entity
type Entity struct {
	Name    string   `json:"name,omitempty"`
	Actions []string `json:"actions,omitempty"`
}
