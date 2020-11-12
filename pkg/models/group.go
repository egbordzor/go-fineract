package models

// Group struct for Group
type Group struct {
	Id                                  int64      `json:"id,omitempty"`
	Office                              Office     `json:"office,omitempty"`
	Staff                               Staff      `json:"staff,omitempty"`
	Parent                              Group      `json:"parent,omitempty"`
	GroupLevel                          GroupLevel `json:"groupLevel,omitempty"`
	GroupMembers                        []Group    `json:"groupMembers,omitempty"`
	ClientMembers                       []Client   `json:"clientMembers,omitempty"`
	SubmittedOnDate                     string     `json:"submittedOnDate,omitempty"`
	AccountNumberRequiresAutoGeneration bool       `json:"accountNumberRequiresAutoGeneration,omitempty"`
	Pending                             bool       `json:"pending,omitempty"`
	Group                               bool       `json:"group,omitempty"`
	NotActive                           bool       `json:"notActive,omitempty"`
	Center                              bool       `json:"center,omitempty"`
	TransferInProgress                  bool       `json:"transferInProgress,omitempty"`
	TransferOnHold                      bool       `json:"transferOnHold,omitempty"`
	NotPending                          bool       `json:"notPending,omitempty"`
	ActivationLocalDate                 string     `json:"activationLocalDate,omitempty"`
	TransferInProgressOrOnHold          bool       `json:"transferInProgressOrOnHold,omitempty"`
	ChildGroup                          bool       `json:"childGroup,omitempty"`
	ActiveClientMembers                 []Client   `json:"activeClientMembers,omitempty"`
	Closed                              bool       `json:"closed,omitempty"`
	Active                              bool       `json:"active,omitempty"`
	New                                 bool       `json:"new,omitempty"`
}
