package models

import (
	"time"
)

// ScorecardValue struct for ScorecardValue
type ScorecardValue struct {
	QuestionId int64     `json:"questionId,omitempty"`
	ResponseId int64     `json:"responseId,omitempty"`
	Value      int32     `json:"value,omitempty"`
	CreatedOn  time.Time `json:"createdOn,omitempty"`
}
