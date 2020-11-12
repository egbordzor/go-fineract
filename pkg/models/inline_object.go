package models

// InlineObject struct for InlineObject
type InlineObject struct {
	File       FormDataContentDisposition `json:"file,omitempty"`
	Locale     string                     `json:"locale,omitempty"`
	DateFormat string                     `json:"dateFormat,omitempty"`
}
