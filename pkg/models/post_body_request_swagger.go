package models

// PostBodyRequestSwagger struct for PostBodyRequestSwagger
type PostBodyRequestSwagger struct {
	OfficeId        int64  `json:"officeId,omitempty"`
	Firstname       string `json:"firstname,omitempty"`
	Lastname        string `json:"lastname,omitempty"`
	ExternalId      string `json:"externalId,omitempty"`
	DateFormat      string `json:"dateFormat,omitempty"`
	Locale          string `json:"locale,omitempty"`
	Active          bool   `json:"active,omitempty"`
	ActivationDate  string `json:"activationDate,omitempty"`
	SubmittedOnDate string `json:"submittedOnDate,omitempty"`
}
