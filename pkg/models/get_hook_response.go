package models

// GetHookResponse GetHookResponse
type GetHookResponse struct {
	Id           int64   `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	DisplayName  string  `json:"displayName,omitempty"`
	IsActive     bool    `json:"isActive,omitempty"`
	CreatedAt    string  `json:"createdAt,omitempty"`
	UpdatedAt    string  `json:"updatedAt,omitempty"`
	TemplateId   int64   `json:"templateId,omitempty"`
	TemplateName string  `json:"templateName,omitempty"`
	Events       []Event `json:"events,omitempty"`
	Config       []Field `json:"config,omitempty"`
}
