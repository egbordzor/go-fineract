package models

import (
	"time"
)

// SurveyData struct for SurveyData
type SurveyData struct {
	Id             int64           `json:"id,omitempty"`
	ComponentDatas []ComponentData `json:"componentDatas,omitempty"`
	QuestionDatas  []QuestionData  `json:"questionDatas,omitempty"`
	Key            string          `json:"key,omitempty"`
	Name           string          `json:"name,omitempty"`
	Description    string          `json:"description,omitempty"`
	CountryCode    string          `json:"countryCode,omitempty"`
	ValidFrom      time.Time       `json:"validFrom,omitempty"`
	ValidTo        time.Time       `json:"validTo,omitempty"`
}
