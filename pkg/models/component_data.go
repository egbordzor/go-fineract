package models

// ComponentData struct for ComponentData
type ComponentData struct {
	Id          int64  `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	Text        string `json:"text,omitempty"`
	Description string `json:"description,omitempty"`
	SequenceNo  int32  `json:"sequenceNo,omitempty"`
}
