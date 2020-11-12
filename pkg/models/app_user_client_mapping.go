package models

// AppUserClientMapping struct for AppUserClientMapping
type AppUserClientMapping struct {
	Id     int64  `json:"id,omitempty"`
	Client Client `json:"client,omitempty"`
	New    bool   `json:"new,omitempty"`
}
