package models

// GetSelfClientsResponse GetSelfClientsResponse
type GetSelfClientsResponse struct {
	TotalFilteredRecords int32                     `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetSelfClientsPageItems `json:"pageItems,omitempty"`
}
