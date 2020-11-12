package models

// GetFloatingRatesRatePeriods struct for GetFloatingRatesRatePeriods
type GetFloatingRatesRatePeriods struct {
	Id                              int32   `json:"id,omitempty"`
	FromDate                        string  `json:"fromDate,omitempty"`
	InterestRate                    float64 `json:"interestRate,omitempty"`
	IsDifferentialToBaseLendingRate bool    `json:"isDifferentialToBaseLendingRate,omitempty"`
	IsActive                        bool    `json:"isActive,omitempty"`
	CreatedBy                       string  `json:"createdBy,omitempty"`
	CreatedOn                       string  `json:"createdOn,omitempty"`
	ModifiedBy                      string  `json:"modifiedBy,omitempty"`
	ModifiedOn                      string  `json:"modifiedOn,omitempty"`
}
