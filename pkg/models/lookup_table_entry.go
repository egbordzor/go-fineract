package models

// LookupTableEntry struct for LookupTableEntry
type LookupTableEntry struct {
	ValueFrom int32   `json:"valueFrom,omitempty"`
	ValueTo   int32   `json:"valueTo,omitempty"`
	Score     float64 `json:"score,omitempty"`
}
