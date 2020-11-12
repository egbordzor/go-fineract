package models

// GetInterestRateChartsChartIdChartSlabsIncentives struct for GetInterestRateChartsChartIdChartSlabsIncentives
type GetInterestRateChartsChartIdChartSlabsIncentives struct {
	Id                 int32                                               `json:"id,omitempty"`
	EntityType         GetInterestRateChartsChartIdChartSlabsEntityType    `json:"entityType,omitempty"`
	AttributeName      GetInterestRateChartsChartIdChartSlabsAttributeName `json:"attributeName,omitempty"`
	ConditionType      GetInterestRateChartsChartIdChartSlabsConditionType `json:"conditionType,omitempty"`
	AttributeValue     int32                                               `json:"attributeValue,omitempty"`
	AttributeValueDesc string                                              `json:"attributeValueDesc,omitempty"`
	IncentiveType      GetInterestRateChartsChartIdChartSlabsIncentiveType `json:"incentiveType,omitempty"`
	Amount             float32                                             `json:"amount,omitempty"`
}
