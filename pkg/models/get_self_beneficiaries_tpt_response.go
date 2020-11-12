package models

// GetSelfBeneficiariesTptResponse GetSelfBeneficiariesTPTResponse
type GetSelfBeneficiariesTptResponse struct {
	Id            int32                              `json:"id,omitempty"`
	Name          string                             `json:"name,omitempty"`
	OfficeName    string                             `json:"officeName,omitempty"`
	ClientName    string                             `json:"clientName,omitempty"`
	AccountType   GetSelfBeneficiariesAccountOptions `json:"accountType,omitempty"`
	AccountNumber int64                              `json:"accountNumber,omitempty"`
	TransferLimit int32                              `json:"transferLimit,omitempty"`
}
