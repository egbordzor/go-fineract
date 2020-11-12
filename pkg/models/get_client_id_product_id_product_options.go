package models

// GetClientIdProductIdProductOptions struct for GetClientIdProductIdProductOptions
type GetClientIdProductIdProductOptions struct {
	Id                                         int32                                                                    `json:"id,omitempty"`
	Name                                       string                                                                   `json:"name,omitempty"`
	ShortName                                  string                                                                   `json:"shortName,omitempty"`
	Description                                string                                                                   `json:"description,omitempty"`
	Currency                                   GetShareAccountsCurrency                                                 `json:"currency,omitempty"`
	TotalShares                                int32                                                                    `json:"totalShares,omitempty"`
	TotalSharesIssued                          int32                                                                    `json:"totalSharesIssued,omitempty"`
	UnitPrice                                  int32                                                                    `json:"unitPrice,omitempty"`
	ShareCapital                               int32                                                                    `json:"shareCapital,omitempty"`
	MinimumShares                              int32                                                                    `json:"minimumShares,omitempty"`
	NorminalShares                             int32                                                                    `json:"norminalShares,omitempty"`
	MaximumShares                              int32                                                                    `json:"maximumShares,omitempty"`
	MarketPrice                                string                                                                   `json:"marketPrice,omitempty"`
	Charges                                    string                                                                   `json:"charges,omitempty"`
	AllowDividendCalculationForInactiveClients bool                                                                     `json:"allowDividendCalculationForInactiveClients,omitempty"`
	LockinPeriod                               int32                                                                    `json:"lockinPeriod,omitempty"`
	LockinPeriodEnum                           GetShareAccountsClientIdProductIdLockPeriodTypeEnum                      `json:"lockinPeriodEnum,omitempty"`
	MinimumActivePeriod                        int32                                                                    `json:"minimumActivePeriod,omitempty"`
	MinimumActivePeriodForDividendsTypeEnum    GetShareAccountsClientIdProductIdMinimumActivePeriodForDividendsTypeEnum `json:"minimumActivePeriodForDividendsTypeEnum,omitempty"`
	AccountingRule                             GetShareAccountsClientIdProductIdAccountingRule                          `json:"accountingRule,omitempty"`
	AccountingMappings                         GetClientIdProductIdAccountingMappings                                   `json:"accountingMappings,omitempty"`
	CurrencyOptions                            GetShareAccountsCurrency                                                 `json:"currencyOptions,omitempty"`
	ChargeOptions                              GetShareAccountsChargeOptions                                            `json:"chargeOptions,omitempty"`
	MinimumActivePeriodFrequencyTypeOptions    GetClientIdProductIdMinimumActivePeriodFrequencyTypeOptions              `json:"minimumActivePeriodFrequencyTypeOptions,omitempty"`
	LockinPeriodFrequencyTypeOptions           GetClientIdProductIdLockinPeriodFrequencyTypeOptions                     `json:"lockinPeriodFrequencyTypeOptions,omitempty"`
	AccountingMappingOptions                   map[string]interface{}                                                   `json:"accountingMappingOptions,omitempty"`
}
