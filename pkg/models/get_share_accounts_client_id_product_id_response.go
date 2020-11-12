package models

// GetShareAccountsClientIdProductIdResponse GetShareAccountsClientIdProductIdResponse
type GetShareAccountsClientIdProductIdResponse struct {
	ProductOptions []GetClientIdProductIdProductOptions `json:"productOptions,omitempty"`
	ChargeOptions  []GetClientIdProductIdChargeOptions  `json:"chargeOptions,omitempty"`
}
