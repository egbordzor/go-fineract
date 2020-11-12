package models

// SingleDebitOrCreditEntryCommand struct for SingleDebitOrCreditEntryCommand
type SingleDebitOrCreditEntryCommand struct {
	GlAccountId               int64    `json:"glAccountId,omitempty"`
	Amount                    float32  `json:"amount,omitempty"`
	Comments                  string   `json:"comments,omitempty"`
	ParametersPassedInRequest []string `json:"parametersPassedInRequest,omitempty"`
	GlAccountIdChanged        bool     `json:"glAccountIdChanged,omitempty"`
	GlAmountChanged           bool     `json:"glAmountChanged,omitempty"`
	CommentsChanged           bool     `json:"commentsChanged,omitempty"`
}
