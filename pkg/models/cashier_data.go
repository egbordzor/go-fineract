package models

import (
	"time"
)

// CashierData struct for CashierData
type CashierData struct {
	Id          int64     `json:"id,omitempty"`
	TellerId    int64     `json:"tellerId,omitempty"`
	OfficeId    int64     `json:"officeId,omitempty"`
	StaffId     int64     `json:"staffId,omitempty"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `json:"startDate,omitempty"`
	EndDate     time.Time `json:"endDate,omitempty"`
	StartTime   string    `json:"startTime,omitempty"`
	EndTime     string    `json:"endTime,omitempty"`
	OfficeName  string    `json:"officeName,omitempty"`
	TellerName  string    `json:"tellerName,omitempty"`
	StaffName   string    `json:"staffName,omitempty"`
	FullDay     bool      `json:"fullDay,omitempty"`
}
