package models

// PostSelfLoansDatatables struct for PostSelfLoansDatatables
type PostSelfLoansDatatables struct {
	RegisteredTableName string            `json:"registeredTableName,omitempty"`
	Data                PostSelfLoansData `json:"data,omitempty"`
}
