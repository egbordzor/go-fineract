package models

// ClientData struct for ClientData
type ClientData struct {
	Id               int64                  `json:"id,omitempty"`
	ExternalId       string                 `json:"externalId,omitempty"`
	ActivationDate   string                 `json:"activationDate,omitempty"`
	Firstname        string                 `json:"firstname,omitempty"`
	Lastname         string                 `json:"lastname,omitempty"`
	OfficeName       string                 `json:"officeName,omitempty"`
	ImageId          int64                  `json:"imageId,omitempty"`
	ImagePresent     bool                   `json:"imagePresent,omitempty"`
	Timeline         map[string]interface{} `json:"timeline,omitempty"`
	SavingsAccountId int64                  `json:"savingsAccountId,omitempty"`
	IsAddressEnabled bool                   `json:"isAddressEnabled,omitempty"`
	RowIndex         int32                  `json:"rowIndex,omitempty"`
}
