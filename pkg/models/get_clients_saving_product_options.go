package models

// GetClientsSavingProductOptions struct for GetClientsSavingProductOptions
type GetClientsSavingProductOptions struct {
	Id                        int32  `json:"id,omitempty"`
	Name                      string `json:"name,omitempty"`
	WithdrawalFeeForTransfers bool   `json:"withdrawalFeeForTransfers,omitempty"`
	AllowOverdraft            bool   `json:"allowOverdraft,omitempty"`
}
