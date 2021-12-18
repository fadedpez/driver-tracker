package entities

import "github.com/fadedpez/driver-tracker/protos"

type Season struct {
	SeasonYear    string                         `json:"season_year"`
	SeasonRounds  *[]Round                       `json:"season_rounds"`
	SeasonDrivers *[]protos.GrandPrixParticipant `json:"season_drivers"`
}
