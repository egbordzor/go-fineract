package models

// GetAccountTransfersTemplateRefundByTransferFromOffice struct for GetAccountTransfersTemplateRefundByTransferFromOffice
type GetAccountTransfersTemplateRefundByTransferFromOffice struct {
	Id            int32  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	NameDecorated string `json:"nameDecorated,omitempty"`
	ExternalId    int32  `json:"externalId,omitempty"`
	OpeningDate   string `json:"openingDate,omitempty"`
	Hierarchy     string `json:"hierarchy,omitempty"`
}
