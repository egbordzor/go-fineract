package models

// PostAccountNumberFormatsRequest PostAccountNumberFormatsRequest
type PostAccountNumberFormatsRequest struct {
	AccountType int64 `json:"accountType,omitempty"`
	PrefixType  int64 `json:"prefixType,omitempty"`
}
