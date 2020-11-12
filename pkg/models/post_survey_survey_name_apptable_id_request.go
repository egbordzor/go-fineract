package models

import (
	"time"
)

// PostSurveySurveyNameApptableIdRequest PostSurveySurveyNameApptableIdRequest
type PostSurveySurveyNameApptableIdRequest struct {
	PpiHouseholdMembersCdQ1Householdmembers     int64     `json:"ppi_household_members_cd_q1_householdmembers,omitempty"`
	PpiHighestschoolCdQ2Highestschool           int64     `json:"ppi_highestschool_cd_q2_highestschool,omitempty"`
	PpiBusinessoccupationCdQ3Businessoccupation int64     `json:"ppi_businessoccupation_cd_q3_businessoccupation,omitempty"`
	PpiHabitableroomsCdQ4Habitablerooms         int64     `json:"ppi_habitablerooms_cd_q4_habitablerooms,omitempty"`
	PpiFloortypeCdQ5Floortype                   int64     `json:"ppi_floortype_cd_q5_floortype,omitempty"`
	PpiLightingsourceCdQ6Lightingsource         int64     `json:"ppi_lightingsource_cd_q6_lightingsource,omitempty"`
	PpiIronsCdQ7Irons                           int64     `json:"ppi_irons_cd_q7_irons,omitempty"`
	PpiMosquitonetsCdQ8Mosquitonets             int64     `json:"ppi_mosquitonets_cd_q8_mosquitonets,omitempty"`
	PpiTowelsCdQ9Towels                         int64     `json:"ppi_towels_cd_q9_towels,omitempty"`
	PpiFryingpansCdQ10Fryingpans                int64     `json:"ppi_fryingpans_cd_q10_fryingpans,omitempty"`
	Date                                        time.Time `json:"Date,omitempty"`
	DateFormat                                  time.Time `json:"dateFormat,omitempty"`
	Locale                                      string    `json:"locale,omitempty"`
}
