package models

// GetStandingInstructionsResponse GetStandingInstructionsResponse
type GetStandingInstructionsResponse struct {
	TotalFilteredRecords int32                                    `json:"totalFilteredRecords,omitempty"`
	PageItems            []GetPageItemsStandingInstructionSwagger `json:"pageItems,omitempty"`
}
