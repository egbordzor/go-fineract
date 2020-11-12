package models

// PostAdhocQuerySearchRequest PostAdhocQuerySearchRequest
type PostAdhocQuerySearchRequest struct {
	Locale                               string `json:"locale,omitempty"`
	DateFormat                           string `json:"dateFormat,omitempty"`
	LoanDateOption                       string `json:"loanDateOption,omitempty"`
	LoanFromDate                         string `json:"loanFromDate,omitempty"`
	LoanToDate                           string `json:"loanToDate,omitempty"`
	IncludeOutStandingAmountPercentage   bool   `json:"includeOutStandingAmountPercentage,omitempty"`
	OutStandingAmountPercentageCondition string `json:"outStandingAmountPercentageCondition,omitempty"`
	OutStandingAmountPercentage          int64  `json:"outStandingAmountPercentage,omitempty"`
	IncludeOutstandingAmount             bool   `json:"includeOutstandingAmount,omitempty"`
	OutstandingAmountCondition           string `json:"outstandingAmountCondition,omitempty"`
	MinOutstandingAmount                 int64  `json:"minOutstandingAmount,omitempty"`
	MaxOutstandingAmount                 int64  `json:"maxOutstandingAmount,omitempty"`
}
