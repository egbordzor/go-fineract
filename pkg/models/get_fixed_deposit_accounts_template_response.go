package models

// GetFixedDepositAccountsTemplateResponse GetFixedDepositAccountsTemplateResponse
type GetFixedDepositAccountsTemplateResponse struct {
	ClientId       int32                                   `json:"clientId,omitempty"`
	ClientName     string                                  `json:"clientName,omitempty"`
	ProductOptions []GetFixedDepositAccountsProductOptions `json:"productOptions,omitempty"`
}
