package models

// BatchResponse struct for BatchResponse
type BatchResponse struct {
	RequestId  int64    `json:"requestId,omitempty"`
	StatusCode int32    `json:"statusCode,omitempty"`
	Headers    []Header `json:"headers,omitempty"`
	Body       string   `json:"body,omitempty"`
}
