package models

import (
	"time"
)

// GetResourceTypeResourceIdNotesNoteIdResponse GetResourceTypeResourceIdNotesNoteIdResponse
type GetResourceTypeResourceIdNotesNoteIdResponse struct {
	Id                int32            `json:"id,omitempty"`
	ClientId          int32            `json:"clientId,omitempty"`
	NoteType          GetNotesNoteType `json:"noteType,omitempty"`
	Note              string           `json:"note,omitempty"`
	CreatedById       int32            `json:"createdById,omitempty"`
	CreatedByUsername string           `json:"createdByUsername,omitempty"`
	CreatedOn         time.Time        `json:"createdOn,omitempty"`
	UpdatedById       int32            `json:"updatedById,omitempty"`
	UpdatedByUsername string           `json:"updatedByUsername,omitempty"`
	UpdatedOn         time.Time        `json:"updatedOn,omitempty"`
}
