package models

// GetSelfClientsClientIdResponse GetSelfClientsClientIdResponse
type GetSelfClientsClientIdResponse struct {
	Id                 int32                  `json:"id,omitempty"`
	AccountNo          int64                  `json:"accountNo,omitempty"`
	Status             GetSelfClientsStatus   `json:"status,omitempty"`
	Active             bool                   `json:"active,omitempty"`
	ActivationDate     string                 `json:"activationDate,omitempty"`
	Firstname          string                 `json:"firstname,omitempty"`
	Lastname           string                 `json:"lastname,omitempty"`
	DisplayName        string                 `json:"displayName,omitempty"`
	OfficeId           int32                  `json:"officeId,omitempty"`
	OfficeName         string                 `json:"officeName,omitempty"`
	Timeline           GetSelfClientsTimeline `json:"timeline,omitempty"`
	SavingsProductId   int32                  `json:"savingsProductId,omitempty"`
	SavingsProductName string                 `json:"savingsProductName,omitempty"`
	Groups             []string               `json:"groups,omitempty"`
}
