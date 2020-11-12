package models

// GetAccountTransfersResponse GetAccountTransfersResponse
type GetAccountTransfersResponse struct {
	TotalFilteredRecords int32                          `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetAccountTransfersPageItems `json:"pageItems,omitempty"`
}
