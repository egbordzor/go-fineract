package models

// PutRecurringDepositAccountsAccountIdResponse PutRecurringDepositAccountsAccountIdResponse
type PutRecurringDepositAccountsAccountIdResponse struct {
	OfficeId   int32                              `json:"officeId,omitempty"`
	ClientId   int32                              `json:"clientId,omitempty"`
	SavingsId  int32                              `json:"savingsId,omitempty"`
	ResourceId int32                              `json:"resourceId,omitempty"`
	Changes    PutRecurringDepositAccountsChanges `json:"changes,omitempty"`
}
