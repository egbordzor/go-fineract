package models

// PostSavingsAccountsRequest PostSavingsAccountsRequest
type PostSavingsAccountsRequest struct {
	ClientId        int32  `json:"clientId,omitempty"`
	ProductId       int32  `json:"productId,omitempty"`
	Locale          string `json:"locale,omitempty"`
	DateFormat      string `json:"dateFormat,omitempty"`
	SubmittedOnDate string `json:"submittedOnDate,omitempty"`
}
