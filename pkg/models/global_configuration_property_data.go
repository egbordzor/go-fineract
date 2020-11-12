package models

import (
	"time"
)

// GlobalConfigurationPropertyData struct for GlobalConfigurationPropertyData
type GlobalConfigurationPropertyData struct {
	Name        string    `json:"name,omitempty"`
	Enabled     bool      `json:"enabled,omitempty"`
	Value       int64     `json:"value,omitempty"`
	DateValue   time.Time `json:"dateValue,omitempty"`
	Id          int64     `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	TrapDoor    bool      `json:"trapDoor,omitempty"`
}
