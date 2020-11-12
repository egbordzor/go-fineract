package models

// GetAccountsTypeAccountIdResponse GetAccountsTypeAccountIdResponse
type GetAccountsTypeAccountIdResponse struct {
	Id                                         int32                         `json:"id,omitempty"`
	AccountNo                                  int64                         `json:"accountNo,omitempty"`
	SavingsAccountNumber                       int64                         `json:"savingsAccountNumber,omitempty"`
	ClientId                                   int32                         `json:"clientId,omitempty"`
	ClientName                                 string                        `json:"clientName,omitempty"`
	ProductId                                  int32                         `json:"productId,omitempty"`
	ProductName                                string                        `json:"productName,omitempty"`
	Status                                     GetAccountsStatus             `json:"status,omitempty"`
	Timeline                                   GetAccountsTimeline           `json:"timeline,omitempty"`
	Currency                                   GetAccountsCurrency           `json:"currency,omitempty"`
	Summary                                    GetAccountsSummary            `json:"summary,omitempty"`
	PurchasedShares                            []GetAccountsPurchasedShares  `json:"purchasedShares,omitempty"`
	SavingsAccountId                           int32                         `json:"savingsAccountId,omitempty"`
	CurrentMarketPrice                         int32                         `json:"currentMarketPrice,omitempty"`
	LockinPeriod                               int32                         `json:"lockinPeriod,omitempty"`
	LockPeriodTypeEnum                         GetAccountsLockPeriodTypeEnum `json:"lockPeriodTypeEnum,omitempty"`
	MinimumActivePeriod                        int32                         `json:"minimumActivePeriod,omitempty"`
	MinimumActivePeriodTypeEnum                GetAccountsLockPeriodTypeEnum `json:"minimumActivePeriodTypeEnum,omitempty"`
	AllowDividendCalculationForInactiveClients bool                          `json:"allowDividendCalculationForInactiveClients,omitempty"`
	Charges                                    []GetAccountsCharges          `json:"charges,omitempty"`
	Dividends                                  []string                      `json:"dividends,omitempty"`
}
