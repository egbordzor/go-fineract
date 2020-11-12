package models

// GetClientsTemplateResponse GetClientsTemplateResponse
type GetClientsTemplateResponse struct {
	ActivationDate       string                           `json:"activationDate,omitempty"`
	OfficeId             int32                            `json:"officeId,omitempty"`
	OfficeOptions        []GetClientsOfficeOptions        `json:"officeOptions,omitempty"`
	StaffOptions         []GetClientsStaffOptions         `json:"staffOptions,omitempty"`
	SavingProductOptions []GetClientsSavingProductOptions `json:"savingProductOptions,omitempty"`
	Datatables           []GetClientsDataTables           `json:"datatables,omitempty"`
}
