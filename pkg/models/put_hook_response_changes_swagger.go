package models

// PutHookResponseChangesSwagger struct for PutHookResponseChangesSwagger
type PutHookResponseChangesSwagger struct {
	DisplayName string  `json:"displayName,omitempty"`
	Events      []Event `json:"events,omitempty"`
	Config      []Field `json:"config,omitempty"`
}
