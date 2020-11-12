package models

// PutStaffRequest PutStaffRequest
type PutStaffRequest struct {
	IsLoanOfficer bool   `json:"isLoanOfficer,omitempty"`
	ExternalId    string `json:"externalId,omitempty"`
}
