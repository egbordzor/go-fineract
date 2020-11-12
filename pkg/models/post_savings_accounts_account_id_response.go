package models

// PostSavingsAccountsAccountIdResponse PostSavingsAccountsAccountIdResponse
type PostSavingsAccountsAccountIdResponse struct {
	OfficeId   int32                  `json:"officeId,omitempty"`
	ClientId   int32                  `json:"clientId,omitempty"`
	ResourceId int32                  `json:"resourceId,omitempty"`
	Changes    map[string]interface{} `json:"changes,omitempty"`
}
