package models

import (
	"time"
)

// FormDataContentDisposition struct for FormDataContentDisposition
type FormDataContentDisposition struct {
	Type             string            `json:"type,omitempty"`
	Parameters       map[string]string `json:"parameters,omitempty"`
	FileName         string            `json:"fileName,omitempty"`
	CreationDate     time.Time         `json:"creationDate,omitempty"`
	ModificationDate time.Time         `json:"modificationDate,omitempty"`
	ReadDate         time.Time         `json:"readDate,omitempty"`
	Size             int64             `json:"size,omitempty"`
	Name             string            `json:"name,omitempty"`
}
