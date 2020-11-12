package models

// GetCentersTemplateResponse GetCentersTemplateResponse
type GetCentersTemplateResponse struct {
	Active         bool                      `json:"active,omitempty"`
	ActivationDate string                    `json:"activationDate,omitempty"`
	OfficeId       int32                     `json:"officeId,omitempty"`
	OfficeOptions  []GetCentersOfficeOptions `json:"officeOptions,omitempty"`
	StaffOptions   []GetCentersStaffOptions  `json:"staffOptions,omitempty"`
}
