package models

// PutCodeValueDataResponse PutCodeValueDataResponse
type PutCodeValueDataResponse struct {
	ResourceId int64                      `json:"resourceId,omitempty"`
	Changes    PutCodeValuechangesSwagger `json:"changes,omitempty"`
}
