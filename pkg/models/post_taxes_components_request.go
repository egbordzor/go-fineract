package models

// PostTaxesComponentsRequest PostTaxesComponentsRequest
type PostTaxesComponentsRequest struct {
	Name              string  `json:"name,omitempty"`
	Percentage        float32 `json:"percentage,omitempty"`
	CreditAccountType int32   `json:"creditAccountType,omitempty"`
	CreditAcountId    int32   `json:"creditAcountId,omitempty"`
	Locale            string  `json:"locale,omitempty"`
	DateFormat        string  `json:"dateFormat,omitempty"`
	StartDate         string  `json:"startDate,omitempty"`
}
