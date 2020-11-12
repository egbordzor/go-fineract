package models

// GetLoansTotal struct for GetLoansTotal
type GetLoansTotal struct {
	CurrencyCode       string  `json:"currencyCode,omitempty"`
	DigitsAfterDecimal int32   `json:"digitsAfterDecimal,omitempty"`
	InMultiplesOf      int32   `json:"inMultiplesOf,omitempty"`
	Amount             float32 `json:"amount,omitempty"`
	DefaultName        string  `json:"defaultName,omitempty"`
	NameCode           string  `json:"nameCode,omitempty"`
	DisplaySymbol      string  `json:"displaySymbol,omitempty"`
	Zero               bool    `json:"zero,omitempty"`
	GreaterThanZero    bool    `json:"greaterThanZero,omitempty"`
	DisplaySymbolValue string  `json:"displaySymbolValue,omitempty"`
}
