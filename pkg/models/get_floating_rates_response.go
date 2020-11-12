package models

// GetFloatingRatesResponse GetFloatingRatesResponse
type GetFloatingRatesResponse struct {
	Id                int32  `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	IsBaseLendingRate bool   `json:"isBaseLendingRate,omitempty"`
	IsActive          bool   `json:"isActive,omitempty"`
	CreatedBy         string `json:"createdBy,omitempty"`
	CreatedOn         string `json:"createdOn,omitempty"`
	ModifiedBy        string `json:"modifiedBy,omitempty"`
	ModifiedOn        string `json:"modifiedOn,omitempty"`
}
