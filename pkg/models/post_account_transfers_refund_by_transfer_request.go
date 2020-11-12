package models

// PostAccountTransfersRefundByTransferRequest PostAccountTransfersRefundByTransferRequest
type PostAccountTransfersRefundByTransferRequest struct {
	FromAccountId       int32   `json:"fromAccountId,omitempty"`
	FromAccountType     int32   `json:"fromAccountType,omitempty"`
	ToOfficeId          int32   `json:"toOfficeId,omitempty"`
	ToClientId          int32   `json:"toClientId,omitempty"`
	ToAccountType       int32   `json:"toAccountType,omitempty"`
	ToAccountId         int32   `json:"toAccountId,omitempty"`
	TransferAmount      float32 `json:"transferAmount,omitempty"`
	TransferDate        string  `json:"transferDate,omitempty"`
	TransferDescription string  `json:"transferDescription,omitempty"`
	Locale              string  `json:"locale,omitempty"`
	DateFormat          string  `json:"dateFormat,omitempty"`
	FromClientId        int32   `json:"fromClientId,omitempty"`
	FromOfficeId        int32   `json:"fromOfficeId,omitempty"`
}
