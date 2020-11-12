package models

// GetLoanWriteOffAccount struct for GetLoanWriteOffAccount
type GetLoanWriteOffAccount struct {
	Id     int32  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	GlCode int32  `json:"glCode,omitempty"`
}
