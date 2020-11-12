package models

// GetAssetType struct for GetAssetType
type GetAssetType struct {
	Id          int32  `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}
