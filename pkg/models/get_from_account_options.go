package models

// GetFromAccountOptions struct for GetFromAccountOptions
type GetFromAccountOptions struct {
	AccountId   int32             `json:"accountId,omitempty"`
	AccountNo   int32             `json:"accountNo,omitempty"`
	AccountType GetAccountOptions `json:"accountType,omitempty"`
	ClientId    int32             `json:"clientId,omitempty"`
	ClientName  string            `json:"clientName,omitempty"`
	OfficeId    int32             `json:"officeId,omitempty"`
	OfficeName  string            `json:"officeName,omitempty"`
}
