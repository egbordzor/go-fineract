package models

// GetSurveyResponse GetSurveyResponse
type GetSurveyResponse struct {
	DatatableData GetSurveyResponseDatatableData `json:"datatableData,omitempty"`
	Enabled       bool                           `json:"enabled,omitempty"`
}
