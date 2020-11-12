package models

// GetSurveyResponseDatatableData struct for GetSurveyResponseDatatableData
type GetSurveyResponseDatatableData struct {
	ApplicationTableName string                      `json:"applicationTableName,omitempty"`
	RegisteredTableName  string                      `json:"registeredTableName,omitempty"`
	ColumnHeaderData     []ResultsetColumnHeaderData `json:"columnHeaderData,omitempty"`
}
