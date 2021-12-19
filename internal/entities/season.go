package entities

import "github.com/fadedpez/driver-tracker/protos"

// A collection of rounds belong to a season
// A collection of drivers belong to a season
//TODO:  A collection of teams belong to a season
//TODO: Drivers should have a total points tally from grand prix results
//TODO: Teams should have a total points tally from their drivers' grand prix results

type Season struct {
	SeasonYear    string                         `json:"season_year"`
	SeasonRounds  *[]Round                       `json:"season_rounds"`
	SeasonDrivers *[]protos.GrandPrixParticipant `json:"season_drivers"`
}

//TODO: Need to see if this does what I intend. It got late and my brain may have stalled.
