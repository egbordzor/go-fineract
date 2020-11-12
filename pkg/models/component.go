package models

// Component struct for Component
type Component struct {
	Id          int64  `json:"id,omitempty"`
	Survey      Survey `json:"survey,omitempty"`
	Key         string `json:"key,omitempty"`
	Text        string `json:"text,omitempty"`
	Description string `json:"description,omitempty"`
	SequenceNo  int32  `json:"sequenceNo,omitempty"`
	New         bool   `json:"new,omitempty"`
}
