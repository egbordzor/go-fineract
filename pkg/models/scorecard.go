package models

import (
	"time"
)

// Scorecard struct for Scorecard
type Scorecard struct {
	Id        int64     `json:"id,omitempty"`
	Survey    Survey    `json:"survey,omitempty"`
	Question  Question  `json:"question,omitempty"`
	Response  Response  `json:"response,omitempty"`
	AppUser   AppUser   `json:"appUser,omitempty"`
	Client    Client    `json:"client,omitempty"`
	CreatedOn time.Time `json:"createdOn,omitempty"`
	Value     int32     `json:"value,omitempty"`
	New       bool      `json:"new,omitempty"`
}
