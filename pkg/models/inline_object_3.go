package models

// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	File FormDataBodyPart `json:"file,omitempty"`
	// name
	Name string `json:"name,omitempty"`
	// description
	Description string `json:"description,omitempty"`
}
