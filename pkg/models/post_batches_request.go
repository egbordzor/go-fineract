package models

// PostBatchesRequest PostBatchesRequest
type PostBatchesRequest struct {
	RequestId   int64                  `json:"requestId,omitempty"`
	RelativeUrl string                 `json:"relativeUrl,omitempty"`
	Method      string                 `json:"method,omitempty"`
	Headers     []Header               `json:"headers,omitempty"`
	Reference   int64                  `json:"reference,omitempty"`
	Body        PostBodyRequestSwagger `json:"body,omitempty"`
}
