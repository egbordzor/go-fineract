package models

// PutProvisioningCriteriaResponse PutProvisioningCriteriaResponse
type PutProvisioningCriteriaResponse struct {
	ResourceId int64                                  `json:"resourceId,omitempty"`
	Changes    PutProvisioningCriteriaResponseChanges `json:"changes,omitempty"`
}
