package models

// GetClientsClientIdTransactionsResponse GetClientsClientIdTransactionsResponse
type GetClientsClientIdTransactionsResponse struct {
	TotalFilteredRecords int32                 `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetClientsPageItems `json:"pageItems,omitempty"`
}
