package models

// GetAccountsTypeTemplateResponse GetAccountsTypeTemplateResponse
type GetAccountsTypeTemplateResponse struct {
	ClientId       int32                           `json:"clientId,omitempty"`
	ClientName     string                          `json:"clientName,omitempty"`
	ProductOptions []GetAccountsTypeProductOptions `json:"productOptions,omitempty"`
}
