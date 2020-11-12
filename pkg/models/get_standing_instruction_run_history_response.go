package models

// GetStandingInstructionRunHistoryResponse GetStandingInstructionRunHistoryResponse
type GetStandingInstructionRunHistoryResponse struct {
	TotalFilteredRecords int32                                            `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetStandingInstructionHistoryPageItemsResponse `json:"pageItems,omitempty"`
}
