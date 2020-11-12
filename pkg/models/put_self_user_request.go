package models

// PutSelfUserRequest PutSelfUserRequest
type PutSelfUserRequest struct {
	Password       string `json:"password,omitempty"`
	RepeatPassword string `json:"repeatPassword,omitempty"`
}
