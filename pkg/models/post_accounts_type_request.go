package models

// PostAccountsTypeRequest PostAccountsTypeRequest
type PostAccountsTypeRequest struct {
	ClientId                                   int32                 `json:"clientId,omitempty"`
	ProductId                                  int32                 `json:"productId,omitempty"`
	RequestedShares                            int32                 `json:"requestedShares,omitempty"`
	ExternalId                                 int32                 `json:"externalId,omitempty"`
	SubmittedDate                              string                `json:"submittedDate,omitempty"`
	MinimumActivePeriod                        int32                 `json:"minimumActivePeriod,omitempty"`
	MinimumActivePeriodFrequencyType           int32                 `json:"minimumActivePeriodFrequencyType,omitempty"`
	LockinPeriodFrequency                      int32                 `json:"lockinPeriodFrequency,omitempty"`
	LockinPeriodFrequencyType                  int32                 `json:"lockinPeriodFrequencyType,omitempty"`
	ApplicationDate                            string                `json:"applicationDate,omitempty"`
	AllowDividendCalculationForInactiveClients bool                  `json:"allowDividendCalculationForInactiveClients,omitempty"`
	Locale                                     string                `json:"locale,omitempty"`
	DateFormat                                 string                `json:"dateFormat,omitempty"`
	Charges                                    []PostAccountsCharges `json:"charges,omitempty"`
	SavingsAccountId                           int32                 `json:"savingsAccountId,omitempty"`
}
