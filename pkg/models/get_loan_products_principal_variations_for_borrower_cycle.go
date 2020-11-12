package models

// GetLoanProductsPrincipalVariationsForBorrowerCycle struct for GetLoanProductsPrincipalVariationsForBorrowerCycle
type GetLoanProductsPrincipalVariationsForBorrowerCycle struct {
	Id                  int32                             `json:"id,omitempty"`
	BorrowerCycleNumber int32                             `json:"borrowerCycleNumber,omitempty"`
	ParamType           GetLoanProductsParamType          `json:"paramType,omitempty"`
	ValueConditionType  GetLoanProductsValueConditionType `json:"valueConditionType,omitempty"`
	MinValue            float32                           `json:"minValue,omitempty"`
	MaxValue            float32                           `json:"maxValue,omitempty"`
	DefaultValue        float32                           `json:"defaultValue,omitempty"`
}
