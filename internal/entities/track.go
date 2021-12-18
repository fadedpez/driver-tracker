package entities

//maybe a bit overboard on what we're putting in the struct but it should be fun

type Track struct {
	Name                 string `json:"track_name"`
	Country              string `json:"track_country"`
	LapDistance          string `json:"track_lap_distance"`
	LapCount             string `json:"track_lap_count"`
	LeftTurnCount        string `json:"track_left_turn_count"`
	RightTurnCount       string `json:"track_right_turn_count"`
	DrsZoneCount         string `json:"track_drs_zone_count"`
	TotalElevationChange string `json:"track_total_elevation_change"`
	ID                   string `json:"track_id"`
}
