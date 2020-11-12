package models

// GetAccountNumberFormatsResponseTemplate GetAccountNumberFormatsResponseTemplate
type GetAccountNumberFormatsResponseTemplate struct {
	AccountTypeOptions []EnumOptionData            `json:"accountTypeOptions,omitempty"`
	PrefixTypeOptions  map[string][]EnumOptionData `json:"prefixTypeOptions,omitempty"`
}
