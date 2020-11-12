package models

// PostAccountTransfersRequest PostAccountTransfersRequest
type PostAccountTransfersRequest struct {
	FromOfficeId        int32   `json:"fromOfficeId,omitempty"`
	FromClientId        int32   `json:"fromClientId,omitempty"`
	FromAccountType     int32   `json:"fromAccountType,omitempty"`
	FromAccountId       int32   `json:"fromAccountId,omitempty"`
	ToOfficeId          int32   `json:"toOfficeId,omitempty"`
	ToClientId          int32   `json:"toClientId,omitempty"`
	ToAccountType       int32   `json:"toAccountType,omitempty"`
	ToAccountId         int32   `json:"toAccountId,omitempty"`
	DateFormat          string  `json:"dateFormat,omitempty"`
	Locale              string  `json:"locale,omitempty"`
	TransferDate        string  `json:"transferDate,omitempty"`
	TransferAmount      float32 `json:"transferAmount,omitempty"`
	TransferDescription string  `json:"transferDescription,omitempty"`
}
