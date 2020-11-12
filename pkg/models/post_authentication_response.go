package models

// PostAuthenticationResponse PostAuthenticationResponse
type PostAuthenticationResponse struct {
	Username                       string         `json:"username,omitempty"`
	UserId                         int64          `json:"userId,omitempty"`
	Base64EncodedAuthenticationKey string         `json:"base64EncodedAuthenticationKey,omitempty"`
	Authenticated                  bool           `json:"authenticated,omitempty"`
	OfficeId                       int64          `json:"officeId,omitempty"`
	OfficeName                     string         `json:"officeName,omitempty"`
	StaffId                        int64          `json:"staffId,omitempty"`
	StaffDisplayName               string         `json:"staffDisplayName,omitempty"`
	OrganisationalRole             EnumOptionData `json:"organisationalRole,omitempty"`
	Roles                          []RoleData     `json:"roles,omitempty"`
	Permissions                    []string       `json:"permissions,omitempty"`
}
