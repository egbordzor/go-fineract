package models

// PostLoansResponse PostLoansResponse
type PostLoansResponse struct {
	Currency                   GetLoansLoanIdCurrency              `json:"currency,omitempty"`
	LoanTermInDays             int32                               `json:"loanTermInDays,omitempty"`
	TotalPrincipalDisbursed    int64                               `json:"totalPrincipalDisbursed,omitempty"`
	TotalPrincipalExpected     int64                               `json:"totalPrincipalExpected,omitempty"`
	TotalPrincipalPaid         int64                               `json:"totalPrincipalPaid,omitempty"`
	TotalInterestCharged       float64                             `json:"totalInterestCharged,omitempty"`
	TotalFeeChargesCharged     int64                               `json:"totalFeeChargesCharged,omitempty"`
	TotalPenaltyChargesCharged int64                               `json:"totalPenaltyChargesCharged,omitempty"`
	TotalWaived                int64                               `json:"totalWaived,omitempty"`
	TotalWrittenOff            int64                               `json:"totalWrittenOff,omitempty"`
	TotalRepaymentExpected     float64                             `json:"totalRepaymentExpected,omitempty"`
	TotalRepayment             int64                               `json:"totalRepayment,omitempty"`
	TotalOutstanding           int64                               `json:"totalOutstanding,omitempty"`
	Periods                    []PostLoansRepaymentSchedulePeriods `json:"periods,omitempty"`
}
