package models

// GetAccountTransfersTemplateRefundByTransferFromAccount struct for GetAccountTransfersTemplateRefundByTransferFromAccount
type GetAccountTransfersTemplateRefundByTransferFromAccount struct {
	Id             int32                                               `json:"id,omitempty"`
	AccountNo      int64                                               `json:"accountNo,omitempty"`
	ClientId       int32                                               `json:"clientId,omitempty"`
	ClientName     string                                              `json:"clientName,omitempty"`
	ProductId      int32                                               `json:"productId,omitempty"`
	ProductName    string                                              `json:"productName,omitempty"`
	FieldOfficerId int32                                               `json:"fieldOfficerId,omitempty"`
	Currency       GetAccountTransfersTemplateRefundByTransferCurrency `json:"currency,omitempty"`
	AmtForTransfer float32                                             `json:"amtForTransfer,omitempty"`
}
