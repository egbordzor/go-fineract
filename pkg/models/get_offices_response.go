package models

// GetOfficesResponse GetOfficesResponse
type GetOfficesResponse struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	NameDecorated string `json:"nameDecorated,omitempty"`
	ExternalId    string `json:"externalId,omitempty"`
	OpeningDate   string `json:"openingDate,omitempty"`
	Hierarchy     string `json:"hierarchy,omitempty"`
}
