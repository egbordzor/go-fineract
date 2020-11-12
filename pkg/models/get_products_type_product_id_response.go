package models

// GetProductsTypeProductIdResponse GetProductsTypeProductIdResponse
type GetProductsTypeProductIdResponse struct {
	Id                                         int32                                                `json:"id,omitempty"`
	Name                                       string                                               `json:"name,omitempty"`
	ShortName                                  string                                               `json:"shortName,omitempty"`
	Description                                string                                               `json:"description,omitempty"`
	Currency                                   GetProductsCurrency                                  `json:"currency,omitempty"`
	TotalShares                                int32                                                `json:"totalShares,omitempty"`
	TotalSharesIssued                          int32                                                `json:"totalSharesIssued,omitempty"`
	UnitPrice                                  int32                                                `json:"unitPrice,omitempty"`
	ShareCapital                               int32                                                `json:"shareCapital,omitempty"`
	MinimumShares                              int32                                                `json:"minimumShares,omitempty"`
	NominalShares                              int32                                                `json:"nominalShares,omitempty"`
	MaximumShares                              int32                                                `json:"maximumShares,omitempty"`
	MarketPrice                                []GetProductsMarketPrice                             `json:"marketPrice,omitempty"`
	Charges                                    []GetProductsCharges                                 `json:"charges,omitempty"`
	AllowDividendCalculationForInactiveClients bool                                                 `json:"allowDividendCalculationForInactiveClients,omitempty"`
	LockinPeriod                               int32                                                `json:"lockinPeriod,omitempty"`
	LockPeriodTypeEnum                         GetLockPeriodTypeEnum                                `json:"lockPeriodTypeEnum,omitempty"`
	MinimumActivePeriod                        int32                                                `json:"minimumActivePeriod,omitempty"`
	MinimumActivePeriodForDividendsTypeEnum    GetLockPeriodTypeEnum                                `json:"minimumActivePeriodForDividendsTypeEnum,omitempty"`
	AccountingRule                             GetProductsAccountingRule                            `json:"accountingRule,omitempty"`
	AccountingMappings                         GetProductsAccountingMappings                        `json:"accountingMappings,omitempty"`
	CurrencyOptions                            []GetChargesCurrency                                 `json:"currencyOptions,omitempty"`
	ChargeOptions                              []GetProductsCharges                                 `json:"chargeOptions,omitempty"`
	MinimumActivePeriodFrequencyTypeOptions    []GetProductsMinimumActivePeriodFrequencyTypeOptions `json:"minimumActivePeriodFrequencyTypeOptions,omitempty"`
	LockinPeriodFrequencyTypeOptions           []GetProductsMinimumActivePeriodFrequencyTypeOptions `json:"lockinPeriodFrequencyTypeOptions,omitempty"`
	AccountingMappingOptions                   GetProductsAccountingMappingOptions                  `json:"accountingMappingOptions,omitempty"`
}
