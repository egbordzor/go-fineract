package models

// PutResourceTypeResourceIdNotesNoteIdResponse PutResourceTypeResourceIdNotesNoteIdResponse
type PutResourceTypeResourceIdNotesNoteIdResponse struct {
	OfficeId   int32           `json:"officeId,omitempty"`
	ClientId   int32           `json:"clientId,omitempty"`
	ResourceId int32           `json:"resourceId,omitempty"`
	Changes    PutNotesChanges `json:"changes,omitempty"`
}
