package models

// GetSavingsAccountsResponse GetSavingsAccountsResponse
type GetSavingsAccountsResponse struct {
	TotalFilteredRecords int32                 `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetSavingsPageItems `json:"pageItems,omitempty"`
}
