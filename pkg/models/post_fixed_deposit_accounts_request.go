package models

// PostFixedDepositAccountsRequest PostFixedDepositAccountsRequest
type PostFixedDepositAccountsRequest struct {
	ClientId                 int32   `json:"clientId,omitempty"`
	ProductId                int32   `json:"productId,omitempty"`
	Locale                   string  `json:"locale,omitempty"`
	DateFormat               string  `json:"dateFormat,omitempty"`
	SubmittedOnDate          string  `json:"submittedOnDate,omitempty"`
	DepositAmount            float32 `json:"depositAmount,omitempty"`
	DepositPeriod            int32   `json:"depositPeriod,omitempty"`
	DepositPeriodFrequencyId int32   `json:"depositPeriodFrequencyId,omitempty"`
}
