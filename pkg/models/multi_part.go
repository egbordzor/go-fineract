package models

// MultiPart struct for MultiPart
type MultiPart struct {
	Entity               map[string]interface{}           `json:"entity,omitempty"`
	Headers              map[string][]string              `json:"headers,omitempty"`
	MediaType            BodyPartMediaType                `json:"mediaType,omitempty"`
	Parent               MultiPart                        `json:"parent,omitempty"`
	Providers            map[string]interface{}           `json:"providers,omitempty"`
	BodyParts            []BodyPart                       `json:"bodyParts,omitempty"`
	ContentDisposition   ContentDisposition               `json:"contentDisposition,omitempty"`
	ParameterizedHeaders map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`
}
