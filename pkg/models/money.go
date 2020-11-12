package models

// Money struct for Money
type Money struct {
	CurrencyCode                string           `json:"currencyCode,omitempty"`
	CurrencyDigitsAfterDecimal  int32            `json:"currencyDigitsAfterDecimal,omitempty"`
	Amount                      float32          `json:"amount,omitempty"`
	GreaterThanZero             bool             `json:"greaterThanZero,omitempty"`
	LessThanZero                bool             `json:"lessThanZero,omitempty"`
	CurrencyInMultiplesOf       int32            `json:"currencyInMultiplesOf,omitempty"`
	AmountDefaultedToNullIfZero float32          `json:"amountDefaultedToNullIfZero,omitempty"`
	Currency                    MonetaryCurrency `json:"currency,omitempty"`
	Zero                        bool             `json:"zero,omitempty"`
}
