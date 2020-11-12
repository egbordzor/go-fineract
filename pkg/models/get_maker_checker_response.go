package models

import (
	"time"
)

// GetMakerCheckerResponse GetMakerCheckerResponse
type GetMakerCheckerResponse struct {
	Id               int64     `json:"id,omitempty"`
	ActionName       string    `json:"actionName,omitempty"`
	EntityName       string    `json:"entityName,omitempty"`
	ResourceId       int64     `json:"resourceId,omitempty"`
	SubresourceId    int64     `json:"subresourceId,omitempty"`
	Maker            string    `json:"maker,omitempty"`
	MadeOnDate       time.Time `json:"madeOnDate,omitempty"`
	Checker          string    `json:"checker,omitempty"`
	CheckedOnDate    time.Time `json:"checkedOnDate,omitempty"`
	ProcessingResult string    `json:"processingResult,omitempty"`
	CommandAsJson    string    `json:"commandAsJson,omitempty"`
	OfficeName       string    `json:"officeName,omitempty"`
	GroupLevelName   string    `json:"groupLevelName,omitempty"`
	GroupName        string    `json:"groupName,omitempty"`
	ClientName       string    `json:"clientName,omitempty"`
	LoanAccountNo    string    `json:"loanAccountNo,omitempty"`
	SavingsAccountNo string    `json:"savingsAccountNo,omitempty"`
	ClientId         int64     `json:"clientId,omitempty"`
	LoanId           int64     `json:"loanId,omitempty"`
	Url              string    `json:"url,omitempty"`
}
