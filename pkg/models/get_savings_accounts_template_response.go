package models

// GetSavingsAccountsTemplateResponse GetSavingsAccountsTemplateResponse
type GetSavingsAccountsTemplateResponse struct {
	ClientId       int32                      `json:"clientId,omitempty"`
	ClientName     string                     `json:"clientName,omitempty"`
	ProductOptions []GetSavingsProductOptions `json:"productOptions,omitempty"`
}
