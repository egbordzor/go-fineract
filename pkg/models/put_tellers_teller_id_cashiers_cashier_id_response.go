package models

// PutTellersTellerIdCashiersCashierIdResponse PutTellersTellerIdCashiersCashierIdResponse
type PutTellersTellerIdCashiersCashierIdResponse struct {
	ResourceId    int64                                              `json:"resourceId,omitempty"`
	SubResourceId int64                                              `json:"subResourceId,omitempty"`
	Changes       PutTellersTellerIdCashiersCashierIdResponseChanges `json:"changes,omitempty"`
}
