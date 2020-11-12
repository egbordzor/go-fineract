package models

// GetStandingInstructionHistoryPageItemsResponse struct for GetStandingInstructionHistoryPageItemsResponse
type GetStandingInstructionHistoryPageItemsResponse struct {
	StandingInstructionId int64                                            `json:"standingInstructionId,omitempty"`
	Name                  string                                           `json:"name,omitempty"`
	FromOffice            GetFromOfficeStandingInstructionSwagger          `json:"fromOffice,omitempty"`
	FromClient            GetStandingInstructionHistoryPageItemsFromClient `json:"fromClient,omitempty"`
	FromAccountType       GetFromAccountTypeStandingInstructionSwagger     `json:"fromAccountType,omitempty"`
	FromAccount           GetStandingInstructionHistoryFromAccount         `json:"fromAccount,omitempty"`
	ToAccountType         GetToAccountTypeStandingInstructionSwagger       `json:"toAccountType,omitempty"`
	ToAccount             GetStandingInstructionHistoryToAccount           `json:"toAccount,omitempty"`
	ToOffice              GetToOfficeStandingInstructionSwagger            `json:"toOffice,omitempty"`
	ToClient              GetStandingInstructionHistoryToClient            `json:"toClient,omitempty"`
	Amount                float32                                          `json:"amount,omitempty"`
	Status                string                                           `json:"status,omitempty"`
	ExecutionTime         string                                           `json:"executionTime,omitempty"`
	ErrorLog              string                                           `json:"errorLog,omitempty"`
}
