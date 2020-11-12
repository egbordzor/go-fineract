package models

// GetSelfSavingsStatus struct for GetSelfSavingsStatus
type GetSelfSavingsStatus struct {
	Id                          int32  `json:"id,omitempty"`
	Code                        string `json:"code,omitempty"`
	Description                 string `json:"description,omitempty"`
	SubmittedAndPendingApproval bool   `json:"submittedAndPendingApproval,omitempty"`
	Approved                    bool   `json:"approved,omitempty"`
	Rejected                    bool   `json:"rejected,omitempty"`
	WithdrawnByApplicant        bool   `json:"withdrawnByApplicant,omitempty"`
	Active                      bool   `json:"active,omitempty"`
	Closed                      bool   `json:"closed,omitempty"`
}
