package models

// GetProvisioningCriteriaCriteriaIdResponse GetProvisioningCriteriaCriteriaIdResponse
type GetProvisioningCriteriaCriteriaIdResponse struct {
	CriteriaId           int64                                `json:"criteriaId,omitempty"`
	CriteriaName         string                               `json:"criteriaName,omitempty"`
	CreatedBy            string                               `json:"createdBy,omitempty"`
	LoanProducts         []LoanProductData                    `json:"loanProducts,omitempty"`
	Provisioningcriteria []ProvisioningCriteriaDefinitionData `json:"provisioningcriteria,omitempty"`
}
