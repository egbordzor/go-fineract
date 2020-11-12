package models

// PutAccountsChanges struct for PutAccountsChanges
type PutAccountsChanges struct {
	DateFormat      string `json:"dateFormat,omitempty"`
	ApplicationDate string `json:"applicationDate,omitempty"`
	RequestedShares int32  `json:"requestedShares,omitempty"`
	Locale          string `json:"locale,omitempty"`
}
