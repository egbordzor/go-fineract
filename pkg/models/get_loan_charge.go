package models

// GetLoanCharge struct for GetLoanCharge
type GetLoanCharge struct {
	Id      int32  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Active  bool   `json:"active,omitempty"`
	Penalty bool   `json:"penalty,omitempty"`
}
