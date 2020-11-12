package models

// GetMakerCheckersSearchTemplateResponse GetMakerCheckersSearchTemplateResponse
type GetMakerCheckersSearchTemplateResponse struct {
	AppUsers          []AppUserData            `json:"appUsers,omitempty"`
	ActionNames       []string                 `json:"actionNames,omitempty"`
	EntityNames       []string                 `json:"entityNames,omitempty"`
	ProcessingResults []ProcessingResultLookup `json:"processingResults,omitempty"`
}
