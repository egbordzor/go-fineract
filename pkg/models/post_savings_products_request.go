package models

// PostSavingsProductsRequest PostSavingsProductsRequest
type PostSavingsProductsRequest struct {
	Name                              string               `json:"name,omitempty"`
	ShortName                         string               `json:"shortName,omitempty"`
	Description                       string               `json:"description,omitempty"`
	CurrencyCode                      string               `json:"currencyCode,omitempty"`
	DigitsAfterDecimal                int32                `json:"digitsAfterDecimal,omitempty"`
	InMultiplesOf                     int32                `json:"inMultiplesOf,omitempty"`
	Locale                            string               `json:"locale,omitempty"`
	NominalAnnualInterestRate         float64              `json:"nominalAnnualInterestRate,omitempty"`
	InterestCompoundingPeriodType     int32                `json:"interestCompoundingPeriodType,omitempty"`
	InterestPostingPeriodType         int32                `json:"interestPostingPeriodType,omitempty"`
	InterestCalculationType           int32                `json:"interestCalculationType,omitempty"`
	InterestCalculationDaysInYearType int32                `json:"interestCalculationDaysInYearType,omitempty"`
	AccountingRule                    int32                `json:"accountingRule,omitempty"`
	Charges                           []PostSavingsCharges `json:"charges,omitempty"`
}
