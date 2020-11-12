package models

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	File FormDataBodyPart `json:"file,omitempty"`
	// name
	Name string `json:"name,omitempty"`
	// description
	Description string `json:"description,omitempty"`
}
