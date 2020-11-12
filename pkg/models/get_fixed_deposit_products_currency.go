package models

// GetFixedDepositProductsCurrency struct for GetFixedDepositProductsCurrency
type GetFixedDepositProductsCurrency struct {
	Code          string `json:"code,omitempty"`
	Name          string `json:"name,omitempty"`
	DecimalPlaces int32  `json:"decimalPlaces,omitempty"`
	InMultiplesOf int32  `json:"inMultiplesOf,omitempty"`
	DisplaySymbol string `json:"displaySymbol,omitempty"`
	NameCode      string `json:"nameCode,omitempty"`
	DisplayLabel  string `json:"displayLabel,omitempty"`
}
