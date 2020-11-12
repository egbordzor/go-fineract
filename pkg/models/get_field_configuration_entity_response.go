package models

// GetFieldConfigurationEntityResponse GetFieldConfigurationEntityResponse
type GetFieldConfigurationEntityResponse struct {
	FieldConfigurationId int32  `json:"fieldConfigurationId,omitempty"`
	Entity               string `json:"entity,omitempty"`
	Subentity            string `json:"subentity,omitempty"`
	Field                string `json:"field,omitempty"`
	IsEnabled            string `json:"is_enabled,omitempty"`
	IsMandatory          string `json:"is_mandatory,omitempty"`
	ValidationRegex      string `json:"validation_regex,omitempty"`
}
