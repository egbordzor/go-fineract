package models

// BodyPartMediaType struct for BodyPartMediaType
type BodyPartMediaType struct {
	Type            string            `json:"type,omitempty"`
	Subtype         string            `json:"subtype,omitempty"`
	Parameters      map[string]string `json:"parameters,omitempty"`
	WildcardType    bool              `json:"wildcardType,omitempty"`
	WildcardSubtype bool              `json:"wildcardSubtype,omitempty"`
}
