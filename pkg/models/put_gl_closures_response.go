package models

// PutGlClosuresResponse PutGlClosuresResponse
type PutGlClosuresResponse struct {
	OfficeId   int64  `json:"officeId,omitempty"`
	ResourceId int64  `json:"resourceId,omitempty"`
	Comments   string `json:"comments,omitempty"`
}
