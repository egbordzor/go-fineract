package models

// GetClientsClientIdChargesResponse GetClientsClientIdChargesResponse
type GetClientsClientIdChargesResponse struct {
	TotalFilteredRecords int32                        `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetClientsChargesPageItems `json:"pageItems,omitempty"`
}
