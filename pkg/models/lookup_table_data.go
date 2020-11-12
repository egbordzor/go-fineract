package models

// LookupTableData struct for LookupTableData
type LookupTableData struct {
	Key         string             `json:"key,omitempty"`
	Description string             `json:"description,omitempty"`
	Entries     []LookupTableEntry `json:"entries,omitempty"`
}
