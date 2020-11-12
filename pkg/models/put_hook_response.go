package models

// PutHookResponse PutHookResponse
type PutHookResponse struct {
	ResourceId int64                         `json:"resourceId,omitempty"`
	Changes    PutHookResponseChangesSwagger `json:"changes,omitempty"`
}
