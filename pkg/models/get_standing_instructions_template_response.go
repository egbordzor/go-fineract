package models

// GetStandingInstructionsTemplateResponse GetStandingInstructionsTemplateResponse
type GetStandingInstructionsTemplateResponse struct {
	FromOffice                 GetFromOfficeResponseStandingInstructionSwagger                   `json:"fromOffice,omitempty"`
	FromAccountType            GetFromAccountTypeResponseStandingInstructionSwagger              `json:"fromAccountType,omitempty"`
	FromOfficeOptions          []GetFromOfficeOptionsResponseStandingInstructionSwagger          `json:"fromOfficeOptions,omitempty"`
	FromClientOptions          []GetFromClientOptionsResponseStandingInstructionSwagger          `json:"fromClientOptions,omitempty"`
	FromAccountTypeOptions     []GetFromAccountTypeOptionsResponseStandingInstructionSwagger     `json:"fromAccountTypeOptions,omitempty"`
	ToOfficeOptions            []GetToOfficeOptionsResponseStandingInstructionSwagger            `json:"toOfficeOptions,omitempty"`
	ToAccountTypeOptions       []GetToAccountTypeOptionsResponseStandingInstructionSwagger       `json:"toAccountTypeOptions,omitempty"`
	TransferTypeOptions        []GetTransferTypeOptionsResponseStandingInstructionSwagger        `json:"transferTypeOptions,omitempty"`
	StatusOptions              []GetStatusOptionsResponseStandingInstructionSwagger              `json:"statusOptions,omitempty"`
	InstructionTypeOptions     []GetInstructionTypeOptionsResponseStandingInstructionSwagger     `json:"instructionTypeOptions,omitempty"`
	PriorityOptions            []GetPriorityOptionsResponseStandingInstructionSwagger            `json:"priorityOptions,omitempty"`
	RecurrenceTypeOptions      []GetRecurrenceTypeOptionsResponseStandingInstructionSwagger      `json:"recurrenceTypeOptions,omitempty"`
	RecurrenceFrequencyOptions []GetRecurrenceFrequencyOptionsResponseStandingInstructionSwagger `json:"recurrenceFrequencyOptions,omitempty"`
}
