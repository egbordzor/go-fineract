package models

// GetEntityDatatableChecksResponse GetEntityDatatableChecksResponse
type GetEntityDatatableChecksResponse struct {
	Id            int64          `json:"id,omitempty"`
	Entity        string         `json:"entity,omitempty"`
	Status        EnumOptionData `json:"status,omitempty"`
	DatatableName string         `json:"datatableName,omitempty"`
	SystemDefined bool           `json:"systemDefined,omitempty"`
	Order         int64          `json:"order,omitempty"`
	ProductId     int64          `json:"productId,omitempty"`
	ProductName   string         `json:"productName,omitempty"`
}
