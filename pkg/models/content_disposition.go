package models

import (
	"time"
)

// ContentDisposition struct for ContentDisposition
type ContentDisposition struct {
	Type             string            `json:"type,omitempty"`
	Parameters       map[string]string `json:"parameters,omitempty"`
	FileName         string            `json:"fileName,omitempty"`
	CreationDate     time.Time         `json:"creationDate,omitempty"`
	ModificationDate time.Time         `json:"modificationDate,omitempty"`
	ReadDate         time.Time         `json:"readDate,omitempty"`
	Size             int64             `json:"size,omitempty"`
}
