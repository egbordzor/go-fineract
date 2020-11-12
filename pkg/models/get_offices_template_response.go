package models

// GetOfficesTemplateResponse GetOfficesTemplateResponse
type GetOfficesTemplateResponse struct {
	OpeningDate    string               `json:"openingDate,omitempty"`
	AllowedParents []GetOfficesResponse `json:"allowedParents,omitempty"`
}
