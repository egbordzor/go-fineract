package models

// LoanProductFloatingRates struct for LoanProductFloatingRates
type LoanProductFloatingRates struct {
	Id                                     int64        `json:"id,omitempty"`
	LoanProduct                            LoanProduct  `json:"loanProduct,omitempty"`
	FloatingRate                           FloatingRate `json:"floatingRate,omitempty"`
	InterestRateDifferential               float32      `json:"interestRateDifferential,omitempty"`
	MinDifferentialLendingRate             float32      `json:"minDifferentialLendingRate,omitempty"`
	DefaultDifferentialLendingRate         float32      `json:"defaultDifferentialLendingRate,omitempty"`
	MaxDifferentialLendingRate             float32      `json:"maxDifferentialLendingRate,omitempty"`
	FloatingInterestRateCalculationAllowed bool         `json:"floatingInterestRateCalculationAllowed,omitempty"`
	New                                    bool         `json:"new,omitempty"`
}
