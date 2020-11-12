package models

// PostCentersCenterIdRequest PostCentersCenterIdRequest
type PostCentersCenterIdRequest struct {
	ClosureReasonId int32  `json:"closureReasonId,omitempty"`
	ClosureDate     string `json:"closureDate,omitempty"`
	Locale          string `json:"locale,omitempty"`
	DateFormat      string `json:"dateFormat,omitempty"`
}
