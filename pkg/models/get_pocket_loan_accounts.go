package models

// GetPocketLoanAccounts struct for GetPocketLoanAccounts
type GetPocketLoanAccounts struct {
	PocketId      int32 `json:"pocketId,omitempty"`
	AccountId     int32 `json:"accountId,omitempty"`
	AccountType   int32 `json:"accountType,omitempty"`
	AccountNumber int32 `json:"accountNumber,omitempty"`
	Id            int32 `json:"id,omitempty"`
}
