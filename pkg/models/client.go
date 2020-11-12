package models

import (
	"time"
)

// Client struct for Client
type Client struct {
	Id                                  int64     `json:"id,omitempty"`
	Office                              Office    `json:"office,omitempty"`
	TransferToOffice                    Office    `json:"transferToOffice,omitempty"`
	Image                               Image     `json:"image,omitempty"`
	Status                              int32     `json:"status,omitempty"`
	Firstname                           string    `json:"firstname,omitempty"`
	Middlename                          string    `json:"middlename,omitempty"`
	Lastname                            string    `json:"lastname,omitempty"`
	DisplayName                         string    `json:"displayName,omitempty"`
	MobileNo                            string    `json:"mobileNo,omitempty"`
	EmailAddress                        string    `json:"emailAddress,omitempty"`
	ExternalId                          string    `json:"externalId,omitempty"`
	Staff                               Staff     `json:"staff,omitempty"`
	Groups                              []Group   `json:"groups,omitempty"`
	AccountNumberRequiresAutoGeneration bool      `json:"accountNumberRequiresAutoGeneration,omitempty"`
	ClosureDate                         string    `json:"closureDate,omitempty"`
	WithdrawalDate                      string    `json:"withdrawalDate,omitempty"`
	SubmittedOnDate                     string    `json:"submittedOnDate,omitempty"`
	LegalForm                           int32     `json:"legalForm,omitempty"`
	ReopenedDate                        string    `json:"reopenedDate,omitempty"`
	ProposedTransferDate                time.Time `json:"proposedTransferDate,omitempty"`
	Rejected                            bool      `json:"rejected,omitempty"`
	Pending                             bool      `json:"pending,omitempty"`
	NotActive                           bool      `json:"notActive,omitempty"`
	TransferInProgress                  bool      `json:"transferInProgress,omitempty"`
	TransferOnHold                      bool      `json:"transferOnHold,omitempty"`
	NotPending                          bool      `json:"notPending,omitempty"`
	ActivationLocalDate                 string    `json:"activationLocalDate,omitempty"`
	TransferInProgressOrOnHold          bool      `json:"transferInProgressOrOnHold,omitempty"`
	OfficeJoiningLocalDate              string    `json:"officeJoiningLocalDate,omitempty"`
	Withdrawn                           bool      `json:"withdrawn,omitempty"`
	NotStaff                            bool      `json:"notStaff,omitempty"`
	RejectedDate                        string    `json:"rejectedDate,omitempty"`
	Closed                              bool      `json:"closed,omitempty"`
	Active                              bool      `json:"active,omitempty"`
	New                                 bool      `json:"new,omitempty"`
}
