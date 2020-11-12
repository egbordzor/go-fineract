package models

// PutFloatingRatesFloatingRateIdResponse PutFloatingRatesFloatingRateIdResponse
type PutFloatingRatesFloatingRateIdResponse struct {
	ResourceId int32                   `json:"resourceId,omitempty"`
	Changes    PutFloatingRatesChanges `json:"changes,omitempty"`
}
