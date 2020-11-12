package models

// GetSelfClientsClientIdChargesResponse GetSelfClientsClientIdChargesResponse
type GetSelfClientsClientIdChargesResponse struct {
	TotalFilteredRecords int32                            `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetSelfClientsChargesPageItems `json:"pageItems,omitempty"`
}
