package models

// GetClientsResponse GetClientsResponse
type GetClientsResponse struct {
	TotalFilteredRecords int32                         `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetClientsPageItemsResponse `json:"pageItems,omitempty"`
}
