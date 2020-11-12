package models

// PostLoansRepaymentSchedulePeriods struct for PostLoansRepaymentSchedulePeriods
type PostLoansRepaymentSchedulePeriods struct {
	Period                          int32  `json:"period,omitempty"`
	DueDate                         string `json:"dueDate,omitempty"`
	PrincipalDisbursed              int64  `json:"principalDisbursed,omitempty"`
	PrincipalLoanBalanceOutstanding int64  `json:"principalLoanBalanceOutstanding,omitempty"`
	FeeChargesDue                   int64  `json:"feeChargesDue,omitempty"`
	FeeChargesOutstanding           int64  `json:"feeChargesOutstanding,omitempty"`
	TotalOriginalDueForPeriod       int64  `json:"totalOriginalDueForPeriod,omitempty"`
	TotalDueForPeriod               int64  `json:"totalDueForPeriod,omitempty"`
	TotalOutstandingForPeriod       int64  `json:"totalOutstandingForPeriod,omitempty"`
	TotalOverdue                    int64  `json:"totalOverdue,omitempty"`
	TotalActualCostOfLoanForPeriod  int64  `json:"totalActualCostOfLoanForPeriod,omitempty"`
}
