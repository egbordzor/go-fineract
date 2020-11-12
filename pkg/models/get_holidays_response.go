package models

// GetHolidaysResponse GetHolidaysResponse
type GetHolidaysResponse struct {
	Id                      int64          `json:"id,omitempty"`
	Name                    string         `json:"name,omitempty"`
	FromDate                string         `json:"fromDate,omitempty"`
	ToDate                  string         `json:"toDate,omitempty"`
	RepaymentsRescheduledTo string         `json:"repaymentsRescheduledTo,omitempty"`
	OfficeId                int64          `json:"officeId,omitempty"`
	Status                  EnumOptionData `json:"status,omitempty"`
}
