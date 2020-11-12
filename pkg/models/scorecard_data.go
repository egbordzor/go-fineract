package models

// ScorecardData struct for ScorecardData
type ScorecardData struct {
	Id              int64            `json:"id,omitempty"`
	UserId          int64            `json:"userId,omitempty"`
	Username        string           `json:"username,omitempty"`
	ClientId        int64            `json:"clientId,omitempty"`
	SurveyId        int64            `json:"surveyId,omitempty"`
	SurveyName      string           `json:"surveyName,omitempty"`
	ScorecardValues []ScorecardValue `json:"scorecardValues,omitempty"`
}
