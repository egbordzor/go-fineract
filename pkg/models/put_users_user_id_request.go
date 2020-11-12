package models

// PutUsersUserIdRequest PutUsersUserIdRequest
type PutUsersUserIdRequest struct {
	Firstname      string `json:"firstname,omitempty"`
	Password       string `json:"password,omitempty"`
	RepeatPassword string `json:"repeatPassword,omitempty"`
}
