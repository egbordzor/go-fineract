package models

// GetSelfClientsClientIdTransactionsTransactionIdResponse GetSelfClientsClientIdTransactionsTransactionIdResponse
type GetSelfClientsClientIdTransactionsTransactionIdResponse struct {
	Id              int32                                  `json:"id,omitempty"`
	OfficeId        int32                                  `json:"officeId,omitempty"`
	OfficeName      string                                 `json:"officeName,omitempty"`
	Type            GetSelfClientsClientIdTransactionsType `json:"type,omitempty"`
	Date            string                                 `json:"date,omitempty"`
	Currency        GetSelfClientsSavingsAccountsCurrency  `json:"currency,omitempty"`
	Amount          float32                                `json:"amount,omitempty"`
	SubmittedOnDate string                                 `json:"submittedOnDate,omitempty"`
	Reversed        bool                                   `json:"reversed,omitempty"`
}
