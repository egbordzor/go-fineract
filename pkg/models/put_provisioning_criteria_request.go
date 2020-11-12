package models

// PutProvisioningCriteriaRequest PutProvisioningCriteriaRequest
type PutProvisioningCriteriaRequest struct {
	CriteriaName         string                               `json:"criteriaName,omitempty"`
	LoanProducts         []LoanProductData                    `json:"loanProducts,omitempty"`
	Provisioningcriteria []ProvisioningCriteriaDefinitionData `json:"provisioningcriteria,omitempty"`
}
