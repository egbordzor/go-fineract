package models

// PutSelfUserResponse PutSelfUserResponse
type PutSelfUserResponse struct {
	OfficeId   int32              `json:"officeId,omitempty"`
	ResourceId int32              `json:"resourceId,omitempty"`
	Changes    PutSelfUserChanges `json:"changes,omitempty"`
}
