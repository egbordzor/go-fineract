package models

// PutHolidaysHolidayIdResponse PutHolidaysHolidayIdResponse
type PutHolidaysHolidayIdResponse struct {
	ResourceId int64                               `json:"resourceId,omitempty"`
	Changes    PutHolidaysHolidayIdResponseChanges `json:"changes,omitempty"`
}
