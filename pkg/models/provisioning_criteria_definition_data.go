package models

// ProvisioningCriteriaDefinitionData struct for ProvisioningCriteriaDefinitionData
type ProvisioningCriteriaDefinitionData struct {
	Id                     int64   `json:"id,omitempty"`
	CategoryId             int64   `json:"categoryId,omitempty"`
	CategoryName           string  `json:"categoryName,omitempty"`
	MinAge                 int64   `json:"minAge,omitempty"`
	MaxAge                 int64   `json:"maxAge,omitempty"`
	ProvisioningPercentage float32 `json:"provisioningPercentage,omitempty"`
	LiabilityAccount       int64   `json:"liabilityAccount,omitempty"`
	LiabilityCode          string  `json:"liabilityCode,omitempty"`
	ExpenseAccount         int64   `json:"expenseAccount,omitempty"`
	ExpenseCode            string  `json:"expenseCode,omitempty"`
}
