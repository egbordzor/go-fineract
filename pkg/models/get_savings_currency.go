package models

// GetSavingsCurrency struct for GetSavingsCurrency
type GetSavingsCurrency struct {
	Code          string `json:"code,omitempty"`
	Name          string `json:"name,omitempty"`
	DecimalPlaces int32  `json:"decimalPlaces,omitempty"`
	DisplaySymbol string `json:"displaySymbol,omitempty"`
	NameCode      string `json:"nameCode,omitempty"`
	DisplayLabel  string `json:"displayLabel,omitempty"`
}
