package entities

import (
	"time"

	"github.com/fadedpez/driver-tracker/protos"
)

//The goal here is to pull in all drivers that participated in the GrandPrix. It is possible for this list to be different than the Qualifying drivers (not common).
// A set of drivers belongs to a grand prix
// Teams belong to a grand prix too? I dunno man!
//TODO: A Grand Prix should have result showing driver finishing classification, total times, fatest lap, and points earned.

type GrandPrix struct {
	GrandPrixName         string                         `json:"grand_prix_name"`
	GrandPrixTrack        string                         `json:"grand_prix_track"`
	GrandPrixDate         *time.Time                     `json:"grand_prix_date"`
	GrandPrixParticipants *[]protos.GrandPrixParticipant `json:"grand_prix_participants"`
	ID                    string                         `json:"id"`
}

//TODO: Need to see if this does what I intend.
