package models

// GetFixedDepositAccountsStatus struct for GetFixedDepositAccountsStatus
type GetFixedDepositAccountsStatus struct {
	Id                          int32  `json:"id,omitempty"`
	Code                        string `json:"code,omitempty"`
	Description                 string `json:"description,omitempty"`
	SubmittedAndPendingApproval bool   `json:"submittedAndPendingApproval,omitempty"`
	Approved                    bool   `json:"approved,omitempty"`
	Rejected                    bool   `json:"rejected,omitempty"`
	WithdrawnByApplicant        bool   `json:"withdrawnByApplicant,omitempty"`
	Active                      bool   `json:"active,omitempty"`
	Closed                      bool   `json:"closed,omitempty"`
	PrematureClosed             bool   `json:"prematureClosed,omitempty"`
	TransferInProgress          bool   `json:"transferInProgress,omitempty"`
	TransferOnHold              bool   `json:"transferOnHold,omitempty"`
}
