package models

// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	File       FormDataContentDisposition `json:"file,omitempty"`
	Locale     string                     `json:"locale,omitempty"`
	DateFormat string                     `json:"dateFormat,omitempty"`
}
