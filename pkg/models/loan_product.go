package models

// LoanProduct struct for LoanProduct
type LoanProduct struct {
	Id                                           int64                                   `json:"id,omitempty"`
	ShortName                                    string                                  `json:"shortName,omitempty"`
	Rates                                        []Rate                                  `json:"rates,omitempty"`
	LoanProductRelatedDetail                     LoanProductRelatedDetail                `json:"loanProductRelatedDetail,omitempty"`
	IncludeInBorrowerCycle                       bool                                    `json:"includeInBorrowerCycle,omitempty"`
	StartDate                                    string                                  `json:"startDate,omitempty"`
	CloseDate                                    string                                  `json:"closeDate,omitempty"`
	ExternalId                                   string                                  `json:"externalId,omitempty"`
	MinimumDaysBetweenDisbursalAndFirstRepayment int32                                   `json:"minimumDaysBetweenDisbursalAndFirstRepayment,omitempty"`
	ProductInterestRecalculationDetails          LoanProductInterestRecalculationDetails `json:"productInterestRecalculationDetails,omitempty"`
	LoanProductGuaranteeDetails                  LoanProductGuaranteeDetails             `json:"loanProductGuaranteeDetails,omitempty"`
	PrincipalThresholdForLastInstallment         float32                                 `json:"principalThresholdForLastInstallment,omitempty"`
	InstallmentAmountInMultiplesOf               int32                                   `json:"installmentAmountInMultiplesOf,omitempty"`
	FloatingRates                                LoanProductFloatingRates                `json:"floatingRates,omitempty"`
	SyncExpectedWithDisbursementDate             bool                                    `json:"syncExpectedWithDisbursementDate,omitempty"`
	LinkedToFloatingInterestRate                 bool                                    `json:"linkedToFloatingInterestRate,omitempty"`
	EqualAmortization                            bool                                    `json:"equalAmortization,omitempty"`
	InterestRecalculationEnabled                 bool                                    `json:"interestRecalculationEnabled,omitempty"`
	MinNominalInterestRatePerPeriod              float32                                 `json:"minNominalInterestRatePerPeriod,omitempty"`
	MaxNominalInterestRatePerPeriod              float32                                 `json:"maxNominalInterestRatePerPeriod,omitempty"`
	MinNumberOfRepayments                        int32                                   `json:"minNumberOfRepayments,omitempty"`
	MaxNumberOfRepayments                        int32                                   `json:"maxNumberOfRepayments,omitempty"`
	MultiDisburseLoan                            bool                                    `json:"multiDisburseLoan,omitempty"`
	RepaymentStrategy                            LoanTransactionProcessingStrategy       `json:"repaymentStrategy,omitempty"`
	AccountingType                               int32                                   `json:"accountingType,omitempty"`
	LoanProductCharges                           []Charge                                `json:"loanProductCharges,omitempty"`
	LoanProductConfigurableAttributes            LoanProductConfigurableAttributes       `json:"loanProductConfigurableAttributes,omitempty"`
	AccountingDisabled                           bool                                    `json:"accountingDisabled,omitempty"`
	CashBasedAccountingEnabled                   bool                                    `json:"cashBasedAccountingEnabled,omitempty"`
	AccrualBasedAccountingEnabled                bool                                    `json:"accrualBasedAccountingEnabled,omitempty"`
	UpfrontAccrualAccountingEnabled              bool                                    `json:"upfrontAccrualAccountingEnabled,omitempty"`
	PeriodicAccrualAccountingEnabled             bool                                    `json:"periodicAccrualAccountingEnabled,omitempty"`
	PrincipalAmount                              Money                                   `json:"principalAmount,omitempty"`
	MinPrincipalAmount                           Money                                   `json:"minPrincipalAmount,omitempty"`
	MaxPrincipalAmount                           Money                                   `json:"maxPrincipalAmount,omitempty"`
	NominalInterestRatePerPeriod                 float32                                 `json:"nominalInterestRatePerPeriod,omitempty"`
	InterestPeriodFrequencyType                  string                                  `json:"interestPeriodFrequencyType,omitempty"`
	NumberOfRepayments                           int32                                   `json:"numberOfRepayments,omitempty"`
	HoldGuaranteeFundsEnabled                    bool                                    `json:"holdGuaranteeFundsEnabled,omitempty"`
	ArrearsBasedOnOriginalSchedule               bool                                    `json:"arrearsBasedOnOriginalSchedule,omitempty"`
	Currency                                     MonetaryCurrency                        `json:"currency,omitempty"`
	New                                          bool                                    `json:"new,omitempty"`
}
