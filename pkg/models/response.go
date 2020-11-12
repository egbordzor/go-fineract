package models

// Response struct for Response
type Response struct {
	Id         int64    `json:"id,omitempty"`
	Question   Question `json:"question,omitempty"`
	Text       string   `json:"text,omitempty"`
	Value      int32    `json:"value,omitempty"`
	SequenceNo int32    `json:"sequenceNo,omitempty"`
	New        bool     `json:"new,omitempty"`
}
