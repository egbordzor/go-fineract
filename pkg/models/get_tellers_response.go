package models

// GetTellersResponse GetTellersResponse
type GetTellersResponse struct {
	Id              int64  `json:"id,omitempty"`
	OfficeId        int64  `json:"officeId,omitempty"`
	DebitAccountId  int64  `json:"debitAccountId,omitempty"`
	CreditAccountId int64  `json:"creditAccountId,omitempty"`
	Name            string `json:"name,omitempty"`
	StartDate       string `json:"startDate,omitempty"`
	Status          string `json:"status,omitempty"`
	OfficeName      string `json:"officeName,omitempty"`
}
