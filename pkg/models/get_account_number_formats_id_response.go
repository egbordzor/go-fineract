package models

// GetAccountNumberFormatsIdResponse GetAccountNumberFormatsIdResponse
type GetAccountNumberFormatsIdResponse struct {
	Id          int64          `json:"id,omitempty"`
	AccountType EnumOptionData `json:"accountType,omitempty"`
	PrefixType  EnumOptionData `json:"prefixType,omitempty"`
}
