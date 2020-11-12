package models

// GetSelfLoansTemplateResponse GetSelfLoansTemplateResponse
type GetSelfLoansTemplateResponse struct {
	ClientId       int32                        `json:"clientId,omitempty"`
	ClientName     string                       `json:"clientName,omitempty"`
	ClientOfficeId int32                        `json:"clientOfficeId,omitempty"`
	Timeline       GetSelfLoansTimeline         `json:"timeline,omitempty"`
	ProductOptions []GetSelfLoansProductOptions `json:"productOptions,omitempty"`
}
