package models

// GetSavingsProductsTemplateResponse GetSavingsProductsTemplateResponse
type GetSavingsProductsTemplateResponse struct {
	Currency                                 GetSavingsCurrency                                    `json:"currency,omitempty"`
	InterestCompoundingPeriodType            GetSavingsProductsInterestCompoundingPeriodType       `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType                GetSavingsProductsInterestPostingPeriodType           `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType                  GetSavingsProductsInterestCalculationType             `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType        GetSavingsProductsInterestCalculationDaysInYearType   `json:"interestCalculationDaysInYearType,omitempty"`
	AccountingRule                           GetSavingsProductsTemplateAccountingRule              `json:"accountingRule,omitempty"`
	CurrencyOptions                          []GetSavingsCurrency                                  `json:"currencyOptions,omitempty"`
	InterestCompoundingPeriodTypeOptions     []GetSavingsProductsInterestCompoundingPeriodType     `json:"interestCompoundingPeriodTypeOptions,omitempty"`
	InterestPostingPeriodTypeOptions         []GetSavingsProductsInterestPostingPeriodType         `json:"interestPostingPeriodTypeOptions,omitempty"`
	InterestCalculationTypeOptions           []GetSavingsProductsInterestCalculationType           `json:"interestCalculationTypeOptions,omitempty"`
	InterestCalculationDaysInYearTypeOptions []GetSavingsProductsInterestCalculationDaysInYearType `json:"interestCalculationDaysInYearTypeOptions,omitempty"`
	LockinPeriodFrequencyTypeOptions         []GetSavingsProductsLockinPeriodFrequencyTypeOptions  `json:"lockinPeriodFrequencyTypeOptions,omitempty"`
	WithdrawalFeeTypeOptions                 []GetSavingsProductsWithdrawalFeeTypeOptions          `json:"withdrawalFeeTypeOptions,omitempty"`
	PaymentTypeOptions                       []GetSavingsProductsPaymentTypeOptions                `json:"paymentTypeOptions,omitempty"`
	AccountingRuleOptions                    []GetSavingsProductsTemplateAccountingRule            `json:"accountingRuleOptions,omitempty"`
	AccountingMappingOptions                 []GetSavingsProductsAccountingMappingOptions          `json:"accountingMappingOptions,omitempty"`
	ChargeOptions                            []GetSavingsProductsChargeOptions                     `json:"chargeOptions,omitempty"`
}
