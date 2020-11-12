package models

import (
	"time"
)

// Survey struct for Survey
type Survey struct {
	Id          int64       `json:"id,omitempty"`
	Components  []Component `json:"components,omitempty"`
	Questions   []Question  `json:"questions,omitempty"`
	Key         string      `json:"key,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	CountryCode string      `json:"countryCode,omitempty"`
	ValidFrom   time.Time   `json:"validFrom,omitempty"`
	ValidTo     time.Time   `json:"validTo,omitempty"`
	New         bool        `json:"new,omitempty"`
}
