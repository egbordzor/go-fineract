package models

// PostEntityDatatableChecksTemplateRequest PostEntityDatatableChecksTemplateRequest
type PostEntityDatatableChecksTemplateRequest struct {
	Entity        string `json:"entity,omitempty"`
	Status        int64  `json:"status,omitempty"`
	DatatableName string `json:"datatableName,omitempty"`
	ProductId     int64  `json:"productId,omitempty"`
}
