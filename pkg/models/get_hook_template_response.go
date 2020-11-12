package models

// GetHookTemplateResponse GetHookTemplateResponse
type GetHookTemplateResponse struct {
	Templates []HookTemplateData `json:"templates,omitempty"`
	Groupings []Grouping         `json:"groupings,omitempty"`
}
