package models

// PutAccountsTypeAccountIdRequest PutAccountsTypeAccountIdRequest
type PutAccountsTypeAccountIdRequest struct {
	Locale          string `json:"locale,omitempty"`
	DateFormat      string `json:"dateFormat,omitempty"`
	ApplicationDate string `json:"applicationDate,omitempty"`
	RequestedShares int32  `json:"requestedShares,omitempty"`
}
