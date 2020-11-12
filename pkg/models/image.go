package models

// Image struct for Image
type Image struct {
	Id          int64  `json:"id,omitempty"`
	Location    string `json:"location,omitempty"`
	StorageType int32  `json:"storageType,omitempty"`
	New         bool   `json:"new,omitempty"`
}
