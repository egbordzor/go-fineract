package models

// PostSelfBeneficiariesTptRequest PostSelfBeneficiariesTPTRequest
type PostSelfBeneficiariesTptRequest struct {
	Locale        string `json:"locale,omitempty"`
	Name          string `json:"name,omitempty"`
	OfficeName    string `json:"officeName,omitempty"`
	AccountNumber int64  `json:"accountNumber,omitempty"`
	AccountType   int32  `json:"accountType,omitempty"`
	TransferLimit int32  `json:"transferLimit,omitempty"`
}
