package models

// LoanProductBorrowerCycleVariationData struct for LoanProductBorrowerCycleVariationData
type LoanProductBorrowerCycleVariationData struct {
	BorrowerCycleNumber int32   `json:"borrowerCycleNumber,omitempty"`
	ParamType           string  `json:"paramType,omitempty"`
	ValueConditionType  string  `json:"valueConditionType,omitempty"`
	DefaultValue        float32 `json:"defaultValue,omitempty"`
}
