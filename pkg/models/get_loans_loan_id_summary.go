package models

// GetLoansLoanIdSummary struct for GetLoansLoanIdSummary
type GetLoansLoanIdSummary struct {
	Currency                           GetLoansLoanIdCurrency              `json:"currency,omitempty"`
	PrincipalDisbursed                 int64                               `json:"principalDisbursed,omitempty"`
	PrincipalPaid                      int64                               `json:"principalPaid,omitempty"`
	PrincipalWrittenOff                int64                               `json:"principalWrittenOff,omitempty"`
	PrincipalOutstanding               int64                               `json:"principalOutstanding,omitempty"`
	PrincipalOverdue                   float64                             `json:"principalOverdue,omitempty"`
	InterestCharged                    int64                               `json:"interestCharged,omitempty"`
	InterestPaid                       int64                               `json:"interestPaid,omitempty"`
	InterestWaived                     int64                               `json:"interestWaived,omitempty"`
	InterestWrittenOff                 int64                               `json:"interestWrittenOff,omitempty"`
	InterestOutstanding                int64                               `json:"interestOutstanding,omitempty"`
	InterestOverdue                    int64                               `json:"interestOverdue,omitempty"`
	FeeChargesCharged                  int64                               `json:"feeChargesCharged,omitempty"`
	FeeChargesDueAtDisbursementCharged int64                               `json:"feeChargesDueAtDisbursementCharged,omitempty"`
	FeeChargesPaid                     int64                               `json:"feeChargesPaid,omitempty"`
	FeeChargesWaived                   int64                               `json:"feeChargesWaived,omitempty"`
	FeeChargesWrittenOff               int64                               `json:"feeChargesWrittenOff,omitempty"`
	FeeChargesOutstanding              int64                               `json:"feeChargesOutstanding,omitempty"`
	FeeChargesOverdue                  int64                               `json:"feeChargesOverdue,omitempty"`
	PenaltyChargesCharged              int64                               `json:"penaltyChargesCharged,omitempty"`
	PenaltyChargesPaid                 int64                               `json:"penaltyChargesPaid,omitempty"`
	PenaltyChargesWaived               int64                               `json:"penaltyChargesWaived,omitempty"`
	PenaltyChargesWrittenOff           int64                               `json:"penaltyChargesWrittenOff,omitempty"`
	PenaltyChargesOutstanding          int64                               `json:"penaltyChargesOutstanding,omitempty"`
	PenaltyChargesOverdue              int64                               `json:"penaltyChargesOverdue,omitempty"`
	TotalExpectedRepayment             int64                               `json:"totalExpectedRepayment,omitempty"`
	TotalRepayment                     int64                               `json:"totalRepayment,omitempty"`
	TotalExpectedCostOfLoan            int64                               `json:"totalExpectedCostOfLoan,omitempty"`
	TotalCostOfLoan                    int64                               `json:"totalCostOfLoan,omitempty"`
	TotalWaived                        int64                               `json:"totalWaived,omitempty"`
	TotalWrittenOff                    int64                               `json:"totalWrittenOff,omitempty"`
	TotalOutstanding                   int64                               `json:"totalOutstanding,omitempty"`
	TotalOverdue                       float64                             `json:"totalOverdue,omitempty"`
	OverdueSinceDate                   string                              `json:"overdueSinceDate,omitempty"`
	LinkedAccount                      GetLoansLoanIdLinkedAccount         `json:"linkedAccount,omitempty"`
	DisbursementDetails                []GetLoansLoanIdDisbursementDetails `json:"disbursementDetails,omitempty"`
	FixedEmiAmount                     float32                             `json:"fixedEmiAmount,omitempty"`
	MaxOutstandingLoanBalance          int64                               `json:"maxOutstandingLoanBalance,omitempty"`
	CanDisburse                        bool                                `json:"canDisburse,omitempty"`
	EmiAmountVariations                []map[string]interface{}            `json:"emiAmountVariations,omitempty"`
	InArrears                          bool                                `json:"inArrears,omitempty"`
	IsNPA                              bool                                `json:"isNPA,omitempty"`
	OverdueCharges                     []GetLoansLoanIdOverdueCharges      `json:"overdueCharges,omitempty"`
}
