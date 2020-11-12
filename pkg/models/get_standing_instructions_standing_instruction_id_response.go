package models

// GetStandingInstructionsStandingInstructionIdResponse GetStandingInstructionsStandingInstructionIdResponse
type GetStandingInstructionsStandingInstructionIdResponse struct {
	Id                   int64                                            `json:"id,omitempty"`
	AccountDetailId      int64                                            `json:"accountDetailId,omitempty"`
	Name                 string                                           `json:"name,omitempty"`
	FromOffice           GetFromOfficeStandingInstructionSwagger          `json:"fromOffice,omitempty"`
	FromClient           GetFromClientStandingInstructionSwagger          `json:"fromClient,omitempty"`
	FromAccountType      GetFromAccountTypeStandingInstructionSwagger     `json:"fromAccountType,omitempty"`
	FromAccount          GetFromAccountStandingInstructionSwagger         `json:"fromAccount,omitempty"`
	ToOffice             GetToOfficeStandingInstructionSwagger            `json:"toOffice,omitempty"`
	ToClient             GetToClientStandingInstructionSwagger            `json:"toClient,omitempty"`
	ToAccountType        GetToAccountTypeStandingInstructionSwagger       `json:"toAccountType,omitempty"`
	ToAccount            GetToAccountStandingInstructionSwagger           `json:"toAccount,omitempty"`
	TransferType         GetTransferTypeStandingInstructionSwagger        `json:"transferType,omitempty"`
	Priority             GetPriorityStandingInstructionSwagger            `json:"priority,omitempty"`
	InstructionType      GetInstructionTypeStandingInstructionSwagger     `json:"instructionType,omitempty"`
	Status               GetStatusStandingInstructionSwagger              `json:"status,omitempty"`
	Amount               float32                                          `json:"amount,omitempty"`
	ValidFrom            string                                           `json:"validFrom,omitempty"`
	RecurrenceType       GetRecurrenceTypeStandingInstructionSwagger      `json:"recurrenceType,omitempty"`
	RecurrenceFrequency  GetRecurrenceFrequencyStandingInstructionSwagger `json:"recurrenceFrequency,omitempty"`
	RecurrenceInterval   int32                                            `json:"recurrenceInterval,omitempty"`
	RecurrenceOnMonthDay string                                           `json:"recurrenceOnMonthDay,omitempty"`
}
