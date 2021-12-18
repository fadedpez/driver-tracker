package entities

type Round struct {
	RoundTrack        Track      `json:"round_track"`
	RoundQualifying   Qualifying `json:"round_qualifying"`
	RoundGrandPrix    GrandPrix  `json:"round_grand_prix"`
	RoundWeekendDates string     `json:"round_weekend_dates"`
	RoundYear         string     `json:"round_year"`
	RoundID           string     `json:"round_id"`
}
