package models

// Question struct for Question
type Question struct {
	Id           int64      `json:"id,omitempty"`
	Survey       Survey     `json:"survey,omitempty"`
	Responses    []Response `json:"responses,omitempty"`
	ComponentKey string     `json:"componentKey,omitempty"`
	Key          string     `json:"key,omitempty"`
	Text         string     `json:"text,omitempty"`
	Description  string     `json:"description,omitempty"`
	SequenceNo   int32      `json:"sequenceNo,omitempty"`
	New          bool       `json:"new,omitempty"`
}
