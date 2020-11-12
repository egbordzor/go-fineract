package models

// Field struct for Field
type Field struct {
	FieldName  string `json:"fieldName,omitempty"`
	FieldValue string `json:"fieldValue,omitempty"`
	FieldType  string `json:"fieldType,omitempty"`
}
