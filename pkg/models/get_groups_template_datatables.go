package models

// GetGroupsTemplateDatatables struct for GetGroupsTemplateDatatables
type GetGroupsTemplateDatatables struct {
	ApplicationTableName string                              `json:"applicationTableName,omitempty"`
	RegisteredTableName  string                              `json:"registeredTableName,omitempty"`
	ColumnHeaderData     []GetGroupsTemplateColumnHeaderData `json:"columnHeaderData,omitempty"`
}
