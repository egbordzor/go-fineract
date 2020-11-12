package models

// AppUserData struct for AppUserData
type AppUserData struct {
	RowIndex        int32        `json:"rowIndex,omitempty"`
	Clients         []ClientData `json:"clients,omitempty"`
	SelfServiceUser bool         `json:"selfServiceUser,omitempty"`
}
