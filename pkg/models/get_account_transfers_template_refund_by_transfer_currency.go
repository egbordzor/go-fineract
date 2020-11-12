package models

// GetAccountTransfersTemplateRefundByTransferCurrency struct for GetAccountTransfersTemplateRefundByTransferCurrency
type GetAccountTransfersTemplateRefundByTransferCurrency struct {
	Code          string `json:"code,omitempty"`
	Name          string `json:"name,omitempty"`
	DecimalPlaces int32  `json:"decimalPlaces,omitempty"`
	InMultiplesOf int32  `json:"inMultiplesOf,omitempty"`
	DisplaySymbol string `json:"displaySymbol,omitempty"`
	NameCode      string `json:"nameCode,omitempty"`
	DisplayLabel  string `json:"displayLabel,omitempty"`
}
