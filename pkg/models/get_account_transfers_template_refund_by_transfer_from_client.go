package models

// GetAccountTransfersTemplateRefundByTransferFromClient struct for GetAccountTransfersTemplateRefundByTransferFromClient
type GetAccountTransfersTemplateRefundByTransferFromClient struct {
	Id                   int32                       `json:"id,omitempty"`
	AccountNo            int64                       `json:"accountNo,omitempty"`
	Status               GetAccountTransfersStatus   `json:"status,omitempty"`
	Active               bool                        `json:"active,omitempty"`
	ActivationDate       string                      `json:"activationDate,omitempty"`
	Firstname            string                      `json:"firstname,omitempty"`
	Lastname             string                      `json:"lastname,omitempty"`
	DisplayName          string                      `json:"displayName,omitempty"`
	Gender               map[string]interface{}      `json:"gender,omitempty"`
	ClientType           map[string]interface{}      `json:"clientType,omitempty"`
	ClientClassification map[string]interface{}      `json:"clientClassification,omitempty"`
	OfficeId             int32                       `json:"officeId,omitempty"`
	OfficeName           string                      `json:"officeName,omitempty"`
	Timeline             GetAccountTransfersTimeline `json:"timeline,omitempty"`
	Groups               map[string]interface{}      `json:"groups,omitempty"`
}
