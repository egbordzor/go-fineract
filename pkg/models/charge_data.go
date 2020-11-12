package models

// ChargeData struct for ChargeData
type ChargeData struct {
	Id                       int64        `json:"id,omitempty"`
	Name                     string       `json:"name,omitempty"`
	Penalty                  bool         `json:"penalty,omitempty"`
	Currency                 CurrencyData `json:"currency,omitempty"`
	OverdueInstallmentCharge bool         `json:"overdueInstallmentCharge,omitempty"`
}
