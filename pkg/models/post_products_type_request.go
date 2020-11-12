package models

// PostProductsTypeRequest PostProductsTypeRequest
type PostProductsTypeRequest struct {
	Name                                       string                           `json:"name,omitempty"`
	ShortName                                  string                           `json:"shortName,omitempty"`
	Description                                string                           `json:"description,omitempty"`
	CurrencyCode                               string                           `json:"currencyCode,omitempty"`
	DigitsAfterDecimal                         int32                            `json:"digitsAfterDecimal,omitempty"`
	InMultiplesOf                              int32                            `json:"inMultiplesOf,omitempty"`
	Locale                                     string                           `json:"locale,omitempty"`
	TotalShares                                int32                            `json:"totalShares,omitempty"`
	SharesIssued                               int32                            `json:"sharesIssued,omitempty"`
	UnitPrice                                  int32                            `json:"unitPrice,omitempty"`
	MinimumShares                              int32                            `json:"minimumShares,omitempty"`
	NominalShares                              int32                            `json:"nominalShares,omitempty"`
	MaximumShares                              int32                            `json:"maximumShares,omitempty"`
	MinimumActivePeriodForDividends            int32                            `json:"minimumActivePeriodForDividends,omitempty"`
	MinimumactiveperiodFrequencyType           int32                            `json:"minimumactiveperiodFrequencyType,omitempty"`
	LockinPeriodFrequency                      int32                            `json:"lockinPeriodFrequency,omitempty"`
	LockinPeriodFrequencyType                  int32                            `json:"lockinPeriodFrequencyType,omitempty"`
	AllowDividendCalculationForInactiveClients bool                             `json:"allowDividendCalculationForInactiveClients,omitempty"`
	MarketPricePeriods                         []PostProductsMarketPricePeriods `json:"marketPricePeriods,omitempty"`
	ChargesSelected                            []PostProductsChargesSelected    `json:"chargesSelected,omitempty"`
	AccountingRule                             int32                            `json:"accountingRule,omitempty"`
}
