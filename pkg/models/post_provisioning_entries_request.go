package models

// PostProvisioningEntriesRequest PostProvisioningEntriesRequest
type PostProvisioningEntriesRequest struct {
	Date                 string `json:"date,omitempty"`
	Locale               string `json:"locale,omitempty"`
	DateFormat           string `json:"dateFormat,omitempty"`
	Createjournalentries string `json:"createjournalentries,omitempty"`
	Provisioningentry    string `json:"provisioningentry,omitempty"`
	Entries              string `json:"entries,omitempty"`
}
