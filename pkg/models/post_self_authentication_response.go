package models

// PostSelfAuthenticationResponse PostSelfAuthenticationResponse
type PostSelfAuthenticationResponse struct {
	Username                       string                               `json:"username,omitempty"`
	UserId                         int32                                `json:"userId,omitempty"`
	Base64EncodedAuthenticationKey string                               `json:"base64EncodedAuthenticationKey,omitempty"`
	Authenticated                  bool                                 `json:"authenticated,omitempty"`
	OfficeId                       int32                                `json:"officeId,omitempty"`
	OfficeName                     string                               `json:"officeName,omitempty"`
	StaffId                        int32                                `json:"staffId,omitempty"`
	StaffDisplayName               string                               `json:"staffDisplayName,omitempty"`
	OrganisationalRole             GetSelfUserDetailsOrganisationalRole `json:"organisationalRole,omitempty"`
	Roles                          []GetSelfUserDetailsRoles            `json:"roles,omitempty"`
	Permissions                    []string                             `json:"permissions,omitempty"`
	IsSelfServiceUser              bool                                 `json:"isSelfServiceUser,omitempty"`
	Clients                        []int32                              `json:"clients,omitempty"`
}
