package models

// GetAccountRulesTemplateResponse GetAccountRulesTemplateResponse
type GetAccountRulesTemplateResponse struct {
	SystemDefined   bool            `json:"systemDefined,omitempty"`
	AllowedOffices  []OfficeData    `json:"allowedOffices,omitempty"`
	AllowedAccounts []GlAccountData `json:"allowedAccounts,omitempty"`
}
