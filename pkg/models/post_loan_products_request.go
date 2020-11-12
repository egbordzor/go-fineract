package models

// PostLoanProductsRequest PostLoanProductsRequest
type PostLoanProductsRequest struct {
	Name                            string  `json:"name,omitempty"`
	ShortName                       string  `json:"shortName,omitempty"`
	CurrencyCode                    string  `json:"currencyCode,omitempty"`
	Locale                          string  `json:"locale,omitempty"`
	DigitsAfterDecimal              int32   `json:"digitsAfterDecimal,omitempty"`
	InMultiplesOf                   int32   `json:"inMultiplesOf,omitempty"`
	Principal                       float64 `json:"principal,omitempty"`
	NumberOfRepayments              int32   `json:"numberOfRepayments,omitempty"`
	RepaymentEvery                  int32   `json:"repaymentEvery,omitempty"`
	RepaymentFrequencyType          int32   `json:"repaymentFrequencyType,omitempty"`
	TransactionProcessingStrategyId int32   `json:"transactionProcessingStrategyId,omitempty"`
	InterestRatePerPeriod           float64 `json:"interestRatePerPeriod,omitempty"`
	InterestRateFrequencyType       int32   `json:"interestRateFrequencyType,omitempty"`
	AmortizationType                int32   `json:"amortizationType,omitempty"`
	InterestType                    int32   `json:"interestType,omitempty"`
	InterestCalculationPeriodType   int32   `json:"interestCalculationPeriodType,omitempty"`
	DaysInMonthType                 int32   `json:"daysInMonthType,omitempty"`
	DaysInYearType                  int32   `json:"daysInYearType,omitempty"`
	IsInterestRecalculationEnabled  bool    `json:"isInterestRecalculationEnabled,omitempty"`
	AccountingRule                  int32   `json:"accountingRule,omitempty"`
	FundSourceAccountId             int32   `json:"fundSourceAccountId,omitempty"`
	LoanPortfolioAccountId          int32   `json:"loanPortfolioAccountId,omitempty"`
	ReceivableInterestAccountId     int32   `json:"receivableInterestAccountId,omitempty"`
	ReceivableFeeAccountId          int32   `json:"receivableFeeAccountId,omitempty"`
	ReceivablePenaltyAccountId      int32   `json:"receivablePenaltyAccountId,omitempty"`
	InterestOnLoanAccountId         int32   `json:"interestOnLoanAccountId,omitempty"`
	IncomeFromFeeAccountId          int32   `json:"incomeFromFeeAccountId,omitempty"`
	IncomeFromPenaltyAccountId      int32   `json:"incomeFromPenaltyAccountId,omitempty"`
	OverpaymentLiabilityAccountId   int32   `json:"overpaymentLiabilityAccountId,omitempty"`
	WriteOffAccountId               int32   `json:"writeOffAccountId,omitempty"`
}
