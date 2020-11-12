package models

// InteropTransactionTypeData struct for InteropTransactionTypeData
type InteropTransactionTypeData struct {
	Scenario      string `json:"scenario"`
	SubScenario   string `json:"subScenario,omitempty"`
	Initiator     string `json:"initiator"`
	InitiatorType string `json:"initiatorType"`
}
