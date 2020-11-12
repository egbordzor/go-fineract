package models

// GetCachesResponse GetCachesResponse
type GetCachesResponse struct {
	CacheType EnumOptionData `json:"cacheType,omitempty"`
	Enabled   bool           `json:"enabled,omitempty"`
}
