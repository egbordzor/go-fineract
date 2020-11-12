package models

// GetEntityDatatableChecksTemplateResponse GetEntityDatatableChecksTemplateResponse
type GetEntityDatatableChecksTemplateResponse struct {
	Entities            []string                 `json:"entities,omitempty"`
	StatusClient        []map[string]interface{} `json:"statusClient,omitempty"`
	StatusGroup         []map[string]interface{} `json:"statusGroup,omitempty"`
	StatusSavings       []map[string]interface{} `json:"statusSavings,omitempty"`
	StatusLoans         []map[string]interface{} `json:"statusLoans,omitempty"`
	Datatables          []map[string]interface{} `json:"datatables,omitempty"`
	LoanProductDatas    []LoanProductData        `json:"loanProductDatas,omitempty"`
	SavingsProductDatas []SavingsProductData     `json:"savingsProductDatas,omitempty"`
}
