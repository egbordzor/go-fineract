package models

// PostProductsMarketPricePeriods struct for PostProductsMarketPricePeriods
type PostProductsMarketPricePeriods struct {
	Locale     string `json:"locale,omitempty"`
	DateFormat string `json:"dateFormat,omitempty"`
	FromDate   string `json:"fromDate,omitempty"`
	ShareValue int32  `json:"shareValue,omitempty"`
}
