package models

// GetProvisioningCriteriaResponse GetProvisioningCriteriaResponse
type GetProvisioningCriteriaResponse struct {
	CriteriaId   int64  `json:"criteriaId,omitempty"`
	CriteriaName string `json:"criteriaName,omitempty"`
	CreatedBy    string `json:"createdBy,omitempty"`
}
