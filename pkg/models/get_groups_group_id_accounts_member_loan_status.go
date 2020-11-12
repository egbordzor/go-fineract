package models

// GetGroupsGroupIdAccountsMemberLoanStatus struct for GetGroupsGroupIdAccountsMemberLoanStatus
type GetGroupsGroupIdAccountsMemberLoanStatus struct {
	Id                   int32  `json:"id,omitempty"`
	Code                 string `json:"code,omitempty"`
	Description          string `json:"description,omitempty"`
	PendingApproval      bool   `json:"pendingApproval,omitempty"`
	WaitingForDisbursal  bool   `json:"waitingForDisbursal,omitempty"`
	Active               bool   `json:"active,omitempty"`
	ClosedObligationsMet bool   `json:"closedObligationsMet,omitempty"`
	ClosedWrittenOff     bool   `json:"closedWrittenOff,omitempty"`
	ClosedRescheduled    bool   `json:"closedRescheduled,omitempty"`
	Closed               bool   `json:"closed,omitempty"`
	Overpaid             bool   `json:"overpaid,omitempty"`
}
