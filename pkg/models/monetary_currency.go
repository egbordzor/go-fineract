package models

// MonetaryCurrency struct for MonetaryCurrency
type MonetaryCurrency struct {
	Code                  string `json:"code,omitempty"`
	DigitsAfterDecimal    int32  `json:"digitsAfterDecimal,omitempty"`
	CurrencyInMultiplesOf int32  `json:"currencyInMultiplesOf,omitempty"`
}
