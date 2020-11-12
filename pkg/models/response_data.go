package models

// ResponseData struct for ResponseData
type ResponseData struct {
	Id         int64  `json:"id,omitempty"`
	Text       string `json:"text,omitempty"`
	Value      int32  `json:"value,omitempty"`
	SequenceNo int32  `json:"sequenceNo,omitempty"`
}
