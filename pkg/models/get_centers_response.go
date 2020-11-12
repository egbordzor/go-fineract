package models

// GetCentersResponse GetCentersResponse
type GetCentersResponse struct {
	TotalFilteredRecords int32                 `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetCentersPageItems `json:"pageItems,omitempty"`
}
