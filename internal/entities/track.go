package entities

type Track struct {
	TrackName                 string `json:"track_name"`
	TrackCountry              string `json:"track_country"`
	TrackLapDistance          string `json:"track_lap_distance"`
	TrackLapCount             string `json:"track_lap_count"`
	TrackLeftTurnCount        string `json:"track_left_turn_count"`
	TrackRightTurnCount       string `json:"track_right_turn_count"`
	TrackDrsZoneCount         string `json:"track_drs_zone_count"`
	TrackTotalElevationChange string `json:"track_total_elevation_change"`
	TrackID                   string `json:"track_id"`
}
