package models

// PostProvisioningCriteriaRequest PostProvisioningCriteriaRequest
type PostProvisioningCriteriaRequest struct {
	CriteriaName         string                               `json:"criteriaName,omitempty"`
	LoanProducts         []LoanProductData                    `json:"loanProducts,omitempty"`
	Provisioningcriteria []ProvisioningCriteriaDefinitionData `json:"provisioningcriteria,omitempty"`
}
