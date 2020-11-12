package models

// GetSelfClientsClientIdTransactionsResponse GetSelfClientsClientIdTransactionsResponse
type GetSelfClientsClientIdTransactionsResponse struct {
	TotalFilteredRecords int32                                         `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetSelfClientsClientIdTransactionsPageItems `json:"pageItems,omitempty"`
}
