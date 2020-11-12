package models

// PostHookRequest PostHookRequest
type PostHookRequest struct {
	Name        string  `json:"name,omitempty"`
	IsActive    bool    `json:"isActive,omitempty"`
	DisplayName string  `json:"displayName,omitempty"`
	TemplateId  int64   `json:"templateId,omitempty"`
	Events      []Event `json:"events,omitempty"`
	Config      []Field `json:"config,omitempty"`
}
