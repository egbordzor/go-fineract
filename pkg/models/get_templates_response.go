package models

// GetTemplatesResponse GetTemplatesResponse
type GetTemplatesResponse struct {
	Id      int64            `json:"id,omitempty"`
	Name    string           `json:"name,omitempty"`
	Entity  int64            `json:"entity,omitempty"`
	Type    int64            `json:"type,omitempty"`
	Text    string           `json:"text,omitempty"`
	Mappers []TemplateMapper `json:"mappers,omitempty"`
}
