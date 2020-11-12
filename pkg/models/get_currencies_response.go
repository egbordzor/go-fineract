package models

// GetCurrenciesResponse GetCurrenciesResponse
type GetCurrenciesResponse struct {
	SelectedCurrencyOptions []CurrencyData `json:"selectedCurrencyOptions,omitempty"`
	CurrencyOptions         []CurrencyData `json:"currencyOptions,omitempty"`
}
