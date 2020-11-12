package models

// GetAccountsTypeResponse GetAccountsTypeResponse
type GetAccountsTypeResponse struct {
	TotalFilteredRecords int32                  `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetAccountsPageItems `json:"pageItems,omitempty"`
}
