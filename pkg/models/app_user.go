package models

import (
	"time"
)

// AppUser struct for AppUser
type AppUser struct {
	Id                      int64                  `json:"id,omitempty"`
	Email                   string                 `json:"email,omitempty"`
	Username                string                 `json:"username,omitempty"`
	Firstname               string                 `json:"firstname,omitempty"`
	Lastname                string                 `json:"lastname,omitempty"`
	Password                string                 `json:"password,omitempty"`
	AccountNonExpired       bool                   `json:"accountNonExpired,omitempty"`
	AccountNonLocked        bool                   `json:"accountNonLocked,omitempty"`
	CredentialsNonExpired   bool                   `json:"credentialsNonExpired,omitempty"`
	Enabled                 bool                   `json:"enabled,omitempty"`
	Deleted                 bool                   `json:"deleted,omitempty"`
	Office                  Office                 `json:"office,omitempty"`
	Staff                   Staff                  `json:"staff,omitempty"`
	Roles                   []Role                 `json:"roles,omitempty"`
	LastTimePasswordUpdated time.Time              `json:"lastTimePasswordUpdated,omitempty"`
	PasswordNeverExpires    bool                   `json:"passwordNeverExpires,omitempty"`
	AppUserClientMappings   []AppUserClientMapping `json:"appUserClientMappings,omitempty"`
	SelfServiceUser         bool                   `json:"selfServiceUser,omitempty"`
	SystemUser              bool                   `json:"systemUser,omitempty"`
	StaffId                 int64                  `json:"staffId,omitempty"`
	StaffDisplayName        string                 `json:"staffDisplayName,omitempty"`
	NotEnabled              bool                   `json:"notEnabled,omitempty"`
	Authorities             []GrantedAuthority     `json:"authorities,omitempty"`
	New                     bool                   `json:"new,omitempty"`
}
