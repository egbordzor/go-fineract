package models

// QuestionData struct for QuestionData
type QuestionData struct {
	Id            int64          `json:"id,omitempty"`
	ResponseDatas []ResponseData `json:"responseDatas,omitempty"`
	ComponentKey  string         `json:"componentKey,omitempty"`
	Key           string         `json:"key,omitempty"`
	Text          string         `json:"text,omitempty"`
	Description   string         `json:"description,omitempty"`
	SequenceNo    int32          `json:"sequenceNo,omitempty"`
}
