package models

// PostClientsRequest PostClientsRequest
type PostClientsRequest struct {
	OfficeId       int32  `json:"officeId,omitempty"`
	Fullname       string `json:"fullname,omitempty"`
	GroupId        int32  `json:"groupId,omitempty"`
	DateFormat     string `json:"dateFormat,omitempty"`
	Locale         string `json:"locale,omitempty"`
	Active         bool   `json:"active,omitempty"`
	ActivationDate string `json:"activationDate,omitempty"`
}
