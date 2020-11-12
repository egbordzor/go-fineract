package models

// FormDataBodyPart struct for FormDataBodyPart
type FormDataBodyPart struct {
	Entity                     map[string]interface{}           `json:"entity,omitempty"`
	Headers                    map[string][]string              `json:"headers,omitempty"`
	MediaType                  BodyPartMediaType                `json:"mediaType,omitempty"`
	Parent                     MultiPart                        `json:"parent,omitempty"`
	Providers                  map[string]interface{}           `json:"providers,omitempty"`
	Simple                     bool                             `json:"simple,omitempty"`
	FormDataContentDisposition FormDataContentDisposition       `json:"formDataContentDisposition,omitempty"`
	ContentDisposition         ContentDisposition               `json:"contentDisposition,omitempty"`
	Name                       string                           `json:"name,omitempty"`
	Value                      string                           `json:"value,omitempty"`
	ParameterizedHeaders       map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`
}
