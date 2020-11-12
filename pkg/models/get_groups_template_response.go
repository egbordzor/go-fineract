package models

// GetGroupsTemplateResponse GetGroupsTemplateResponse
type GetGroupsTemplateResponse struct {
	OfficeId      int32                            `json:"officeId,omitempty"`
	OfficeOptions []GetGroupsTemplateOfficeOptions `json:"officeOptions,omitempty"`
	StaffOptions  []GetGroupsTemplateStaffOptions  `json:"staffOptions,omitempty"`
	ClientOptions []GetGroupsTemplateClientOptions `json:"clientOptions,omitempty"`
	Datatables    []GetGroupsTemplateDatatables    `json:"datatables,omitempty"`
}
