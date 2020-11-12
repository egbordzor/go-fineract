package models

// GetLoansTemplateResponse GetLoansTemplateResponse
type GetLoansTemplateResponse struct {
	ClientId       int64                            `json:"clientId,omitempty"`
	ClientName     string                           `json:"clientName,omitempty"`
	ClientOfficeId int32                            `json:"clientOfficeId,omitempty"`
	Timeline       GetLoansTemplateTimeline         `json:"timeline,omitempty"`
	ProductOptions []GetLoansTemplateProductOptions `json:"productOptions,omitempty"`
}
