package models

// FinancialActivityData struct for FinancialActivityData
type FinancialActivityData struct {
	Id                  int32  `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	MappedGLAccountType string `json:"mappedGLAccountType,omitempty"`
}
