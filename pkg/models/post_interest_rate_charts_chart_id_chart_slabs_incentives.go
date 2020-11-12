package models

// PostInterestRateChartsChartIdChartSlabsIncentives struct for PostInterestRateChartsChartIdChartSlabsIncentives
type PostInterestRateChartsChartIdChartSlabsIncentives struct {
	EntityType     int32   `json:"entityType,omitempty"`
	AttributeName  int32   `json:"attributeName,omitempty"`
	ConditionType  int32   `json:"conditionType,omitempty"`
	AttributeValue int32   `json:"attributeValue,omitempty"`
	IncentiveType  int32   `json:"incentiveType,omitempty"`
	Amount         float32 `json:"amount,omitempty"`
}
