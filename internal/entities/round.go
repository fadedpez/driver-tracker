package entities

// A track belongs to a Round
// A qualifying session belongs to a Round
// A grand prix (race) belongs to a Round
// A round belongs to a season

type Round struct {
	Track        Track      `json:"round_track"`
	Qualifying   Qualifying `json:"round_qualifying"`
	GrandPrix    GrandPrix  `json:"round_grand_prix"`
	WeekendDates string     `json:"round_weekend_dates"`
	Year         string     `json:"round_year"`
	ID           string     `json:"id"`
}

//TODO: Need to see if this does what I intend.
