package models

// GetGlAccountsTemplateResponse GetGLAccountsTemplateResponse
type GetGlAccountsTemplateResponse struct {
	Disabled                      bool             `json:"disabled,omitempty"`
	ManualEntriesAllowed          bool             `json:"manualEntriesAllowed,omitempty"`
	Type                          EnumOptionData   `json:"type,omitempty"`
	Usage                         EnumOptionData   `json:"usage,omitempty"`
	AccountTypeOptions            []EnumOptionData `json:"accountTypeOptions,omitempty"`
	UsageOptions                  []EnumOptionData `json:"usageOptions,omitempty"`
	AssetHeaderAccountOptions     []GlAccountData  `json:"assetHeaderAccountOptions,omitempty"`
	LiabilityHeaderAccountOptions []GlAccountData  `json:"liabilityHeaderAccountOptions,omitempty"`
	EquityHeaderAccountOptions    []GlAccountData  `json:"equityHeaderAccountOptions,omitempty"`
	ExpenseHeaderAccountOptions   []GlAccountData  `json:"expenseHeaderAccountOptions,omitempty"`
	AllowedAssetsTagOptions       []CodeValueData  `json:"allowedAssetsTagOptions,omitempty"`
	AllowedLiabilitiesTagOptions  []CodeValueData  `json:"allowedLiabilitiesTagOptions,omitempty"`
	AllowedEquityTagOptions       []CodeValueData  `json:"allowedEquityTagOptions,omitempty"`
	AllowedIncomeTagOptions       []CodeValueData  `json:"allowedIncomeTagOptions,omitempty"`
	AllowedExpensesTagOptions     []CodeValueData  `json:"allowedExpensesTagOptions,omitempty"`
}
