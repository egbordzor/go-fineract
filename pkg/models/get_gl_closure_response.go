package models

// GetGlClosureResponse GetGLClosureResponse
type GetGlClosureResponse struct {
	Id                    int64  `json:"id,omitempty"`
	OfficeId              int64  `json:"officeId,omitempty"`
	OfficeName            string `json:"officeName,omitempty"`
	ClosingDate           string `json:"closingDate,omitempty"`
	Deleted               bool   `json:"deleted,omitempty"`
	CreatedDate           string `json:"createdDate,omitempty"`
	LastUpdatedDate       string `json:"lastUpdatedDate,omitempty"`
	CreatedByUserId       int64  `json:"createdByUserId,omitempty"`
	CreatedByUsername     string `json:"createdByUsername,omitempty"`
	LastUpdatedByUserId   int64  `json:"lastUpdatedByUserId,omitempty"`
	LastUpdatedByUsername string `json:"lastUpdatedByUsername,omitempty"`
	Comments              string `json:"comments,omitempty"`
}
