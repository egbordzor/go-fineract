package models

// GetRecurringDepositAccountsTemplateResponse GetRecurringDepositAccountsTemplateResponse
type GetRecurringDepositAccountsTemplateResponse struct {
	ClientId       int32                        `json:"clientId,omitempty"`
	ClientName     string                       `json:"clientName,omitempty"`
	ProductOptions []GetRecurringProductOptions `json:"productOptions,omitempty"`
}
