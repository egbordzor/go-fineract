package models

// GetDataTablesResponse GetDataTablesResponse
type GetDataTablesResponse struct {
	AppTableName  string                      `json:"appTableName,omitempty"`
	DatatableName string                      `json:"datatableName,omitempty"`
	Column        []ResultsetColumnHeaderData `json:"column,omitempty"`
}
